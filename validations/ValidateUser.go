package validations

import "tbo_backend/objects"

func ValidateUser(obj *objects.AuthReq) objects.Response {

	mobileNumber := validateMobile(obj.Mobile)
	if !mobileNumber.Status {
		return mobileNumber
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func ValidateMobileVerify(obj *objects.ValidateOtpReq) objects.Response {

	mobileNumber := validateMobile(obj.Mobile)
	if !mobileNumber.Status {
		return mobileNumber
	}

	otp := validateOtp(obj.Otp)
	if !otp.Status {
		return otp
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func ValidateFollow(obj *objects.FollowReq) objects.Response {

	userId := validateUserId(obj.UserId)
	if !userId.Status {
		return userId
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func ValidateBlockUser(obj *objects.FollowReq) objects.Response {

	userId := validateUserId(obj.UserId)
	if !userId.Status {
		return userId
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func validateUserId(userId string) objects.Response {

	if userId == "" {
		return objects.Response{
			Status: false,
			Msg:    objects.UserIdEmpty,
		}
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func validateMobile(mobile string) objects.Response {

	if mobile == "" || len(mobile) != 10 {
		return objects.Response{
			Status: false,
			Msg:    objects.MobileNumberInvalid,
		}
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func validateOtp(otp string) objects.Response {

	if otp == "" || len(otp) != 4 {
		return objects.Response{
			Status: false,
			Msg:    objects.OtpInvalid,
		}
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func ValidateProfile(obj *objects.ValidateProfileReq) objects.Response {

	name := validateName(obj.Name)
	if !name.Status {
		return name
	}

	username := validateUsername(obj.Username)
	if !username.Status {
		return username
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func validateUsername(username string) objects.Response {

	if username == "" {
		return objects.Response{
			Status: false,
			Msg:    objects.UsernameInvalid,
		}
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}

func validateName(name string) objects.Response {

	if name == "" {
		return objects.Response{
			Status: false,
			Msg:    objects.NameInvalid,
		}
	}

	return objects.Response{
		Status: true,
		Msg:    "",
	}
}
