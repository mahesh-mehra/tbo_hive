package applications

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"
	"sync"
	"tbo_backend/clients"
	"tbo_backend/compute/middlewares"
	"tbo_backend/objects"
	"tbo_backend/repositories"
	"tbo_backend/routes/auth"
	"tbo_backend/routes/profile"
	"tbo_backend/routes/social"
	"tbo_backend/services"
	"tbo_backend/utils"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

// client represents a single WebSocket connection
type client struct{}

var (
	// Map to track active clients.
	// We use a map[*websocket.Conn]client for O(1) lookups/deletes.
	socketClients = make(map[*websocket.Conn]client)

	// Mutex to protect the clients map from concurrent access
	register = sync.Mutex{}
)

func ConnectHttp() bool {

	defer utils.HandlePanic()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("" +
			"<h1>TBO Profiler Links</h1></br>" +
			"<a href='/debug/pprof/'>Profiler Links</a> <br/>" +
			"<a href='/debug/pprof/cmdline'>CMDLine</a> <br/>" +
			"<a href='/debug/pprof/profile'>Profiler</a> <br/>" +
			"<a href='/debug/pprof/symbol'>Symbols</a> <br/>" +
			"<a href='/debug/pprof/trace'>Trace</a> <br/>" +
			"<a href='/debug/pprof/block'>Block</a> <br/>" +
			"<a href='/debug/pprof/goroutine'>Goroutine</a> <br/>" +
			"<a href='/debug/pprof/heap'>Heap</a> <br/>" +
			"<a href='/debug/pprof/threadcreate'>Threadcreate</a> <br/>" +
			"<p>Run the following program</p> </br>" +
			"1. sudo apt-get install graphviz </br>" +
			"2. go tool pprof -web http://server_path:3700/debug/pprof/profile?seconds=10 </br>" +
			"3. go tool pprof -http=:8080 downloaded_file.pb.gz" +
			""))
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	// profiler links
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// hosting the profiler in 3700 port
	go func() {
		err := http.ListenAndServe(objects.ConfigObj.Http, mux)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Profiler started in port :3700")
	}()

	fmt.Println("Starting Http/1.1 server at: ", objects.ConfigObj.Http)

	// hosting the fiber framework
	app = fiber.New(fiber.Config{
		Prefork:      false, // goob.ConfigObj.Worker not found
		ServerHeader: "TBO",
		AppName:      "TBO App v1.0.0",
		ReadTimeout:  18 * time.Second,
	})

	app.Use(func(c *fiber.Ctx) error {

		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "*")
		c.Set("Access-Control-Allow-Headers", "*")

		if strings.ToLower(string(c.Request().Header.Method())) == "options" {
			c.Status(200)
			return nil
		}

		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		c.Next()

		return nil
	})

	// universal apis
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("TBO App")
	})

	// Initialize Dependencies
	userRepo := repositories.NewScyllaUserRepository(clients.ScyllaSession)
	authService := services.NewAuthServiceImpl(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	app.Post("/login", authHandler.Login)
	app.Post("/validate_otp", authHandler.ValidateOtp)

	authG := app.Group("/session", middlewares.UserSessionValidation)

	authG.Post("/update_profile", profile.UpdateProfile)
	authG.Post("/upload_profile_photo", profile.UploadProfilePhoto)
	authG.Get("/fetch_profile", profile.FetchProfile)
	authG.Post("/block_user", social.BlockUser)
	authG.Post("/unblock_user", social.UnBlockUser)
	authG.Get("/fetch_blocked_user_list", social.FetchBlockedUserList)
	// authG.Post("/comment_on_post", social.CommentOnPost)
	authG.Post("/deactivate_account", social.DeactivateAccount)
	authG.Post("/delete_account", social.DeleteAccount)
	authG.Post("/upload_image", social.UploadImage)
	// authG.Post("/upload_video", social.UploadVideo)
	// authG.Post("/share_post", social.SharePost)
	// authG.Post("/report_post", social.ReportPost)
	authG.Post("/follow", social.Follow)
	authG.Post("/unfollow", social.UnFollow)
	// authG.Get("/fetch_stories", social.FetchStories)
	// authG.Get("/fetch_post", social.FetchPost)
	// authG.Post("/fetch_people", social.FetchPeople)
	// authG.Post("/fetch_chat_list", social.FetchChatList)
	// authG.Get("/fetch_trips", trips.FetchTrips)
	// authG.Get("/trip_information", trips.TripInformation)
	// authG.Get("/render360_image_tour", trips.Render360ImageTour)
	// authG.Get("/render_vr_video", trips.RenderVrVideo)

	// 1. Upgrade Middleware
	// This ensures the request is actually a WebSocket upgrade request.
	authG.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// 2. WebSocket Route
	authG.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// --- CONNECT ---
		// Add new client to the map safely
		register.Lock()
		socketClients[c] = client{}
		register.Unlock()

		fmt.Println("New client connected")

		defer func() {
			// --- DISCONNECT ---
			// Remove client from map when function exits
			register.Lock()
			delete(socketClients, c)
			register.Unlock()
			c.Close()
			fmt.Println("Client disconnected")
		}()

		// --- LISTENER LOOP ---
		// Keep reading messages until connection breaks
		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				// Error reading (client closed or network issue)
				break
			}

			fmt.Printf("Received: %s", msg)

			// --- BROADCAST ---
			// Send the received message to ALL other connected clients
			broadcast(messageType, msg)
		}
	}))

	err := app.Listen(objects.ConfigObj.Http)

	if err != nil {
		return false
	}

	fmt.Println("Http/1.1 server started at: ", objects.ConfigObj.Http)

	return true
}

func GetFiberApp() *fiber.App {
	return app
}

// broadcast sends a message to all active clients safely
func broadcast(messageType int, msg []byte) {
	register.Lock()
	defer register.Unlock()

	for conn := range socketClients {
		// Attempt to write the message
		if err := conn.WriteMessage(messageType, msg); err != nil {
			// If write fails, close and remove the bad connection
			conn.Close()
			delete(socketClients, conn)
		}
	}
}
