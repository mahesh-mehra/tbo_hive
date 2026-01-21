package booking_behaviour_model

import (
	"fmt"
	"github.com/korsajan/gocatboost"
	"sync"
	"tbo_backend/objects"
)

var (
	bookingBehaviourModel *gocatboost.Catboost
	loadOnce              sync.Once
	loadErr               error
)

func LoadBookingBehaviourModel() error {
	loadOnce.Do(func() {
		// loading model in memory
		cb, err := gocatboost.FromFile(objects.ConfigObj.CatboostModelsPaths.BookingBehaviourModel)

		// checking for error
		if err != nil {
			loadErr = err
			return
		}

		// updating booking behaviour model
		bookingBehaviourModel = cb

		fmt.Println("Booking Behaviour Model loaded!")
	})

	return loadErr
}
