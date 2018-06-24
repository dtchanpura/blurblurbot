package bot

// Version to be returned in response when asked for
const Version = "0.2.1-rc2"

// BuildDate to be returned in response with version
const BuildDate = "Tue Nov  7 00:15:20 IST 2017"

// HelpText to be returned on /help call.
const HelpText = `Send a photo and get a photo blurred photo background.
There can be atleast two of the following arguments in caption of photo (in given order).
1. Scale: How much bigger the blurred background should be. e.g. 2, 1.5, 5, etc.
2. BlurRadius: Intensity of blurring the scaled backgroud image.
3. SameSizeReturn: By scaling it will increase the image size, if you want to keep image of same size use default "true" as value. e.g. if you have 128x128 photo and scaled 2 the return image will be of 128x128 and the original image will be pasted in center resized to 64x64 which may decrease the quality of image, to avoid that you can set "false" to this argument which can response enlarged image.`
