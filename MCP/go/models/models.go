package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AssetResponseAttributes represents the AssetResponseAttributes schema from the OpenAPI specification
type AssetResponseAttributes struct {
	Url string `json:"url,omitempty"` // The asset file name.
	Created string `json:"created,omitempty"` // The time the asset was created.
	Filename string `json:"filename,omitempty"` // The asset file name.
	Id string `json:"id,omitempty"` // The unique id of the hosted asset in UUID format.
	Region string `json:"region,omitempty"` // The region the asset is hosted, currently only `au` (Australia).
	Renderid string `json:"renderId,omitempty"` // The original render id that created the asset in UUID format. Multiple assets can share the same render id.
	Updated string `json:"updated,omitempty"` // The time the asset status was last updated.
	Owner string `json:"owner,omitempty"` // The owner id of the render task.
	Status string `json:"status,omitempty"` // The status of the asset. <ul> <li>`importing` - the asset is being copied to the hosting service</li> <li>`ready` - the asset is ready to be served to users</li> <li>`failed` - the asset failed to copy or delete</li> <li>`deleted` - the asset has been deleted</li> </ul>
}

// Output represents the Output schema from the OpenAPI specification
type Output struct {
	Quality string `json:"quality,omitempty"` // Adjust the output quality of the video, image or audio. Adjusting quality affects render speed, download speeds and storage requirements due to file size. The default `medium` provides the most optimized choice for all three factors. <ul> <li>`low` - slightly reduced quality, smaller file size</li> <li>`medium` - optimized quality, render speeds and file size</li> <li>`high` - slightly increased quality, larger file size</li> </ul>
	RangeField Range `json:"range,omitempty"` // Specify a time range to render, i.e. to render only a portion of a video or audio file. Omit this setting to export the entire video. Range can also be used to render a frame at a specific time point - setting a range and output format as `jpg` will output a single frame image at the range `start` point.
	Resolution string `json:"resolution,omitempty"` // The output resolution of the video or image. <ul> <li>`preview` - 512px x 288px @ 15fps</li> <li>`mobile` - 640px x 360px @ 25fps</li> <li>`sd` - 1024px x 576px @ 25fps</li> <li>`hd` - 1280px x 720px @ 25fps</li> <li>`1080` - 1920px x 1080px @ 25fps</li> </ul>
	Format string `json:"format"` // The output format and type of media file to generate. <ul> <li>`mp4` - mp4 video file</li> <li>`gif` - animated gif</li> <li>`jpg` - jpg image file</li> <li>`png` - png image file</li> <li>`bmp` - bmp image file</li> <li>`mp3` - mp3 audio file (audio only)</li> </ul>
	Fps int `json:"fps,omitempty"` // Override the default frames per second. Useful for when the source footage is recorded at 30fps, i.e. on mobile devices. Lower frame rates can be used to add cinematic quality (24fps) or to create smaller file size/faster render times or animated gifs (12 or 15fps). Default is 25fps. <ul> <li>`12` - 12fps</li> <li>`15` - 15fps</li> <li>`24` - 24fps</li> <li>`25` - 25fps</li> <li>`30` - 30fps</li> </ul>
	Poster Poster `json:"poster,omitempty"` // Generate a poster image for the video at a specific point from the timeline. The poster image size will match the size of the output video.
	Scaleto string `json:"scaleTo,omitempty"` // Override the resolution and scale the video or image to render at a different size. When using scaleTo the asset should be edited at the resolution dimensions, i.e. use font sizes that look best at HD, then use scaleTo to output the file at SD and the text will be scaled to the correct size. This is useful if you want to create multiple asset sizes. <ul> <li>`preview` - 512px x 288px @ 15fps</li> <li>`mobile` - 640px x 360px @ 25fps</li> <li>`sd` - 1024px x 576px @25fps</li> <li>`hd` - 1280px x 720px @25fps</li> <li>`1080` - 1920px x 1080px @25fps</li> </ul>
	Size Size `json:"size,omitempty"` // Set a custom size for a video or image. When using a custom size omit the `resolution` and `aspectRatio`. Custom sizes must be divisible by 2 based on the encoder specifications.
	Thumbnail Thumbnail `json:"thumbnail,omitempty"` // Generate a thumbnail image for the video or image at a specific point from the timeline.
	Aspectratio string `json:"aspectRatio,omitempty"` // The aspect ratio (shape) of the video or image. Useful for social media output formats. Options are: <ul> <li>`16:9` - regular landscape/horizontal aspect ratio (default)</li> <li>`9:16` - vertical/portrait aspect ratio</li> <li>`1:1` - square aspect ratio</li> <li>`4:5` - short vertical/portrait aspect ratio</li> <li>`4:3` - legacy TV aspect ratio</li> </ul>
	Destinations []ShotstackDestination `json:"destinations,omitempty"` // A destination is a location where output files can be sent to for serving or hosting. By default all rendered assets are automatically sent to the Shotstack hosting destination. [DestinationShotstack](/#tocs_shotstackdestination) is currently the only option with plans to add more in the future such as S3, YouTube, Vimeo and Mux. If you do not require hosting you can opt-out using the `exclude` property.
}

// ShotstackDestination represents the ShotstackDestination schema from the OpenAPI specification
type ShotstackDestination struct {
	Exclude bool `json:"exclude,omitempty"` // Set to `true` to opt-out from the Shotstack hosting and CDN service. All files must be downloaded within 24 hours of rendering.
	Provider string `json:"provider"` // The destination to send rendered assets to - set to `shotstack` for Shotstack hosting and CDN.
}

// AssetRenderResponse represents the AssetRenderResponse schema from the OpenAPI specification
type AssetRenderResponse struct {
	Data []AssetResponseData `json:"data,omitempty"` // An array of asset resources grouped by render id.
}

// Clip represents the Clip schema from the OpenAPI specification
type Clip struct {
	Offset Offset `json:"offset,omitempty"` // Offsets the position of an asset horizontally or vertically by a relative distance.
	Effect string `json:"effect,omitempty"` // A motion effect to apply to the Clip. <ul> <li>`zoomIn` - slow zoom in</li> <li>`zoomOut` - slow zoom out</li> <li>`slideLeft` - slow slide (pan) left</li> <li>`slideRight` - slow slide (pan) right</li> <li>`slideUp` - slow slide (pan) up</li> <li>`slideDown` - slow slide (pan) down</li> </ul>
	Filter string `json:"filter,omitempty"` // A filter effect to apply to the Clip. <ul> <li>`boost` - boost contrast and saturation</li> <li>`contrast` - increase contrast</li> <li>`darken` - darken the scene</li> <li>`greyscale` - remove colour</li> <li>`lighten` - lighten the scene</li> <li>`muted` - reduce saturation and contrast</li> <li>`invert` - invert colors</li> </ul>
	Fit string `json:"fit,omitempty"` // Set how the asset should be scaled to fit the viewport using one of the following options: <ul> <li>`cover` - stretch the asset to fill the viewport without maintaining the aspect ratio.</li> <li>`contain` - fit the entire asset within the viewport while maintaining the original aspect ratio.</li> <li>`crop` - scale the asset to fill the viewport while maintaining the aspect ratio. The asset will be cropped if it exceeds the bounds of the viewport.</li> <li>`none` - preserves the original asset dimensions and does not apply any scaling.</li> </ul>
	Opacity float64 `json:"opacity,omitempty"` // Sets the opacity of the Clip where 1 is opaque and 0 is transparent.
	Scale float64 `json:"scale,omitempty"` // Scale the asset to a fraction of the viewport size - i.e. setting the scale to 0.5 will scale asset to half the size of the viewport. This is useful for picture-in-picture video and scaling images such as logos and watermarks.
	Start float64 `json:"start"` // The start position of the Clip on the timeline, in seconds.
	Position string `json:"position,omitempty"` // Place the asset in one of nine predefined positions of the viewport. This is most effective for when the asset is scaled and you want to position the element to a specific position. <ul> <li>`top` - top (center)</li> <li>`topRight` - top right</li> <li>`right` - right (center)</li> <li>`bottomRight` - bottom right</li> <li>`bottom` - bottom (center)</li> <li>`bottomLeft` - bottom left</li> <li>`left` - left (center)</li> <li>`topLeft` - top left</li> <li>`center` - center</li> </ul>
	Transition Transition `json:"transition,omitempty"` // In and out transitions for a clip - i.e. fade in and fade out
	Asset interface{} `json:"asset"` // The type of asset to display for the duration of this Clip. Value must be one of <b>TitleAsset</b>, <b>ImageAsset</b>, <b>VideoAsset</b>, <b>HtmlAsset</b>, <b>AudioAsset</b> or <b>LumaAsset</b>
	Length float64 `json:"length"` // The length, in seconds, the Clip should play for.
}

// VideoAsset represents the VideoAsset schema from the OpenAPI specification
type VideoAsset struct {
	Crop Crop `json:"crop,omitempty"` // Crop the sides of an asset by a relative amount. The size of the crop is specified using a scale between 0 and 1, relative to the screen width - i.e a left crop of 0.5 will crop half of the asset from the left, a top crop of 0.25 will crop the top by quarter of the asset.
	Src string `json:"src"` // The video source URL. The URL must be publicly accessible or include credentials.
	Trim float64 `json:"trim,omitempty"` // The start trim point of the video clip, in seconds (defaults to 0). Videos will start from the in trim point. The video will play until the file ends or the Clip length is reached.
	TypeField string `json:"type"` // The type of asset - set to `video` for videos.
	Volume float64 `json:"volume,omitempty"` // Set the volume for the video clip between 0 and 1 where 0 is muted and 1 is full volume (defaults to 0).
}

// AudioAsset represents the AudioAsset schema from the OpenAPI specification
type AudioAsset struct {
	Src string `json:"src"` // The audio source URL. The URL must be publicly accessible or include credentials.
	Trim float64 `json:"trim,omitempty"` // The start trim point of the audio clip, in seconds (defaults to 0). Audio will start from the in trim point. The audio will play until the file ends or the Clip length is reached.
	TypeField string `json:"type"` // The type of asset - set to `audio` for audio assets.
	Volume float64 `json:"volume,omitempty"` // Set the volume for the audio clip between 0 and 1 where 0 is muted and 1 is full volume (defaults to 1).
	Effect string `json:"effect,omitempty"` // The effect to apply to the audio asset <ul> <li>`fadeIn` - fade volume in only</li> <li>`fadeOut` - fade volume out only</li> <li>`fadeInFadeOut` - fade volume in and out</li> </ul>
}

// Edit represents the Edit schema from the OpenAPI specification
type Edit struct {
	Output Output `json:"output"` // The output format, render range and type of media to generate.
	Timeline Timeline `json:"timeline"` // A timeline represents the contents of a video edit over time, an audio edit over time, in seconds, or an image layout. A timeline consists of layers called tracks. Tracks are composed of titles, images, audio, html or video segments referred to as clips which are placed along the track at specific starting point and lasting for a specific amount of time.
	Callback string `json:"callback,omitempty"` // An optional webhook callback URL used to receive status notifications when a render completes or fails. See [webhooks](https://shotstack.gitbook.io/docs/guides/architecting-an-application/webhooks) for more details.
	Disk string `json:"disk,omitempty"` // The disk type to use for storing footage and assets for each render. See [disk types](https://shotstack.gitbook.io/docs/guides/architecting-an-application/disk-types) for more details. <ul> <li>`local` - optimized for high speed rendering with up to 512MB storage</li> <li>`mount` - optimized for larger file sizes and longer videos with 5GB for source footage and 512MB for output render</li> </ul>
}

// AssetResponse represents the AssetResponse schema from the OpenAPI specification
type AssetResponse struct {
	Data AssetResponseData `json:"data,omitempty"` // The type of resource (an asset) and attributes of the asset.
}

// QueuedResponse represents the QueuedResponse schema from the OpenAPI specification
type QueuedResponse struct {
	Message string `json:"message"` // `Created`, `Bad Request` or an error message.
	Response QueuedResponseData `json:"response"` // The response data returned with the [QueuedResponse](#tocs_queuedresponse).
	Success bool `json:"success"` // `true` if successfully queued, else `false`.
}

// Soundtrack represents the Soundtrack schema from the OpenAPI specification
type Soundtrack struct {
	Effect string `json:"effect,omitempty"` // The effect to apply to the audio file <ul> <li>`fadeIn` - fade volume in only</li> <li>`fadeOut` - fade volume out only</li> <li>`fadeInFadeOut` - fade volume in and out</li> </ul>
	Src string `json:"src"` // The URL of the mp3 audio file. The URL must be publicly accessible or include credentials.
	Volume float64 `json:"volume,omitempty"` // Set the volume for the soundtrack between 0 and 1 where 0 is muted and 1 is full volume (defaults to 1).
}

// HtmlAsset represents the HtmlAsset schema from the OpenAPI specification
type HtmlAsset struct {
	Width int `json:"width,omitempty"` // Set the width of the HTML asset bounding box in pixels. Text will wrap to fill the bounding box.
	Background string `json:"background,omitempty"` // Apply a background color behind the HTML bounding box using. Set the text color using hexadecimal color notation. Transparency is supported by setting the first two characters of the hex string (opposite to HTML), i.e. #80ffffff will be white with 50% transparency.
	Css string `json:"css,omitempty"` // The CSS text string to apply styling to the HTML. See list of [support CSS properties](https://shotstack.gitbook.io/docs/guides/architecting-an-application/html-support#supported-html-tags).
	Height int `json:"height,omitempty"` // Set the width of the HTML asset bounding box in pixels. Text and elements will be masked if they exceed the height of the bounding box.
	Html string `json:"html"` // The HTML text string. See list of [supported HTML tags](https://shotstack.gitbook.io/docs/guides/architecting-an-application/html-support#supported-html-tags).
	Position string `json:"position,omitempty"` // Place the HTML in one of nine predefined positions within the HTML area. <ul> <li>`top` - top (center)</li> <li>`topRight` - top right</li> <li>`right` - right (center)</li> <li>`bottomRight` - bottom right</li> <li>`bottom` - bottom (center)</li> <li>`bottomLeft` - bottom left</li> <li>`left` - left (center)</li> <li>`topLeft` - top left</li> <li>`center` - center</li> </ul>
	TypeField string `json:"type"` // The type of asset - set to `html` for HTML.
}

// QueuedResponseData represents the QueuedResponseData schema from the OpenAPI specification
type QueuedResponseData struct {
	Message string `json:"message"` // Success response message or error details.
	Id string `json:"id"` // The id of the render task in UUID format.
}

// Offset represents the Offset schema from the OpenAPI specification
type Offset struct {
	X float32 `json:"x,omitempty"` // Offset an asset on the horizontal axis (left or right), range varies from -1 to 1. Positive numbers move the asset right, negative left. For all assets except titles the distance moved is relative to the width of the viewport - i.e. an X offset of 0.5 will move the asset half the screen width to the right.
	Y float32 `json:"y,omitempty"` // Offset an asset on the vertical axis (up or down), range varies from -1 to 1. Positive numbers move the asset up, negative down. For all assets except titles the distance moved is relative to the height of the viewport - i.e. an Y offset of 0.5 will move the asset up half the screen height.
}

// Range represents the Range schema from the OpenAPI specification
type Range struct {
	Length float32 `json:"length,omitempty"` // The length of the portion of the video or audio to render - i.e. render 6 seconds of the video.
	Start float32 `json:"start,omitempty"` // The point on the timeline, in seconds, to start the render from - i.e. start at second 3.
}

// Poster represents the Poster schema from the OpenAPI specification
type Poster struct {
	Capture float64 `json:"capture"` // The point on the timeline in seconds to capture a single frame to use as the poster image.
}

// Size represents the Size schema from the OpenAPI specification
type Size struct {
	Height int `json:"height,omitempty"` // Set a custom height for the video or image file. Value must be divisible by 2. Maximum video height is 1920px, maximum image height is 4096px.
	Width int `json:"width,omitempty"` // Set a custom width for the video or image file. Value must be divisible by 2. Maximum video width is 1920px, maximum image width is 4096px.
}

// Transition represents the Transition schema from the OpenAPI specification
type Transition struct {
	Out string `json:"out,omitempty"` // The transition out. Available transitions are: <ul> <li>`fade` - fade out</li> <li>`reveal` - reveal from right to left</li> <li>`wipeLeft` - fade across screen to the left</li> <li>`wipeRight` - fade across screen to the right</li> <li>`slideLeft` - move slightly left and fade out</li> <li>`slideRight` - move slightly right and fade out</li> <li>`slideUp` - move slightly up and fade out</li> <li>`slideDown` - move slightly down and fade out</li> <li>`carouselLeft` - slide out from right to left</li> <li>`carouselRight` - slide out from left to right</li> <li>`carouselUp` - slide out from bottom to top</li> <li>`carouselDown` - slide out from top to bottom</li> <li>`shuffleTopRight` - rotate out from top right</li> <li>`shuffleRightTop` - rotate out from right top</li> <li>`shuffleRightBottom` - rotate out from right bottom</li> <li>`shuffleBottomRight` - rotate out from bottom right</li> <li>`shuffleBottomLeft` - rotate out from bottom left</li> <li>`shuffleLeftBottom` - rotate out from left bottom</li> <li>`shuffleLeftTop` - rotate out from left top</li> <li>`shuffleTopLeft` - rotate out from top left</li> <li>`zoom` - fast zoom out</li> </ul> The transition speed can also be controlled by appending `Fast` or `Slow` to the transition, e.g. `fadeFast` or `CarouselLeftSlow`.
	In string `json:"in,omitempty"` // The transition in. Available transitions are: <ul> <li>`fade` - fade in</li> <li>`reveal` - reveal from left to right</li> <li>`wipeLeft` - fade across screen to the left</li> <li>`wipeRight` - fade across screen to the right</li> <li>`slideLeft` - move slightly left and fade in</li> <li>`slideRight` - move slightly right and fade in</li> <li>`slideUp` - move slightly up and fade in</li> <li>`slideDown` - move slightly down and fade in</li> <li>`carouselLeft` - slide in from right to left</li> <li>`carouselRight` - slide in from left to right</li> <li>`carouselUp` - slide in from bottom to top</li> <li>`carouselDown` - slide in from top to bottom</li> <li>`shuffleTopRight` - rotate in from top right</li> <li>`shuffleRightTop` - rotate in from right top</li> <li>`shuffleRightBottom` - rotate in from right bottom</li> <li>`shuffleBottomRight` - rotate in from bottom right</li> <li>`shuffleBottomLeft` - rotate in from bottom left</li> <li>`shuffleLeftBottom` - rotate in from left bottom</li> <li>`shuffleLeftTop` - rotate in from left top</li> <li>`shuffleTopLeft` - rotate in from top left</li> <li>`zoom` - fast zoom in</li> </ul> The transition speed can also be controlled by appending `Fast` or `Slow` to the transition, e.g. `fadeFast` or `CarouselLeftSlow`.
}

// TitleAsset represents the TitleAsset schema from the OpenAPI specification
type TitleAsset struct {
	Background string `json:"background,omitempty"` // Apply a background color behind the text. Set the text color using hexadecimal color notation. Transparency is supported by setting the first two characters of the hex string (opposite to HTML), i.e. #80ffffff will be white with 50% transparency. Omit to use transparent background.
	Color string `json:"color,omitempty"` // Set the text color using hexadecimal color notation. Transparency is supported by setting the first two characters of the hex string (opposite to HTML), i.e. #80ffffff will be white with 50% transparency.
	Offset Offset `json:"offset,omitempty"` // Offsets the position of an asset horizontally or vertically by a relative distance.
	Position string `json:"position,omitempty"` // Place the title in one of nine predefined positions of the viewport. <ul> <li>`top` - top (center)</li> <li>`topRight` - top right</li> <li>`right` - right (center)</li> <li>`bottomRight` - bottom right</li> <li>`bottom` - bottom (center)</li> <li>`bottomLeft` - bottom left</li> <li>`left` - left (center)</li> <li>`topLeft` - top left</li> <li>`center` - center</li> </ul>
	Size string `json:"size,omitempty"` // Set the relative size of the text using predefined sizes from xx-small to xx-large. <ul> <li>`xx-small`</li> <li>`x-small`</li> <li>`small`</li> <li>`medium`</li> <li>`large`</li> <li>`x-large`</li> <li>`xx-large`</li> </ul>
	Style string `json:"style,omitempty"` // Uses a preset to apply font properties and styling to the title. <ul> <li>`minimal`</li> <li>`blockbuster`</li> <li>`vogue`</li> <li>`sketchy`</li> <li>`skinny`</li> <li>`chunk`</li> <li>`chunkLight`</li> <li>`marker`</li> <li>`future`</li> <li>`subtitle`</li> </ul>
	Text string `json:"text"` // The title text string - i.e. "My Title".
	TypeField string `json:"type"` // The type of asset - set to `title` for titles.
}

// Font represents the Font schema from the OpenAPI specification
type Font struct {
	Src string `json:"src"` // The URL of the font file. The URL must be publicly accessible or include credentials.
}

// Crop represents the Crop schema from the OpenAPI specification
type Crop struct {
	Left float32 `json:"left,omitempty"` // Crop from the left of the asset
	Right float32 `json:"right,omitempty"` // Crop from the left of the asset
	Top float32 `json:"top,omitempty"` // Crop from the top of the asset
	Bottom float32 `json:"bottom,omitempty"` // Crop from the bottom of the asset
}

// LumaAsset represents the LumaAsset schema from the OpenAPI specification
type LumaAsset struct {
	Src string `json:"src"` // The luma matte source URL. The URL must be publicly accessible or include credentials.
	Trim float64 `json:"trim,omitempty"` // The start trim point of the luma matte clip, in seconds (defaults to 0). Videos will start from the in trim point. A luma matte video will play until the file ends or the Clip length is reached.
	TypeField string `json:"type"` // The type of asset - set to `luma` for luma mattes.
}

// AssetResponseData represents the AssetResponseData schema from the OpenAPI specification
type AssetResponseData struct {
	TypeField string `json:"type,omitempty"` // The type of resource, in this case it is an assets.
	Attributes AssetResponseAttributes `json:"attributes,omitempty"` // The list of asset attributes and their values.
}

// Thumbnail represents the Thumbnail schema from the OpenAPI specification
type Thumbnail struct {
	Capture float64 `json:"capture"` // The point on the timeline in seconds to capture a single frame to use as the thumbnail image.
	Scale float64 `json:"scale"` // Scale the thumbnail size to a fraction of the viewport size - i.e. setting the scale to 0.5 will scale the thumbnail to half the size of the viewport.
}

// Track represents the Track schema from the OpenAPI specification
type Track struct {
	Clips []Clip `json:"clips"` // An array of Clips comprising of TitleClip, ImageClip or VideoClip.
}

// RenderResponseData represents the RenderResponseData schema from the OpenAPI specification
type RenderResponseData struct {
	Url string `json:"url,omitempty"` // The URL of the final asset. This will only be available if status is done. This is a temporary URL and will be deleted after 24 hours. By default all assets are copied to the Shotstack hosting and CDN destination.
	Data Edit `json:"data"` // An edit defines the arrangement of a video on a timeline, an audio edit or an image design and the output format.
	ErrorField string `json:"error,omitempty"` // An error message, only displayed if an error occurred.
	Id string `json:"id"` // The id of the render task in UUID format.
	Poster string `json:"poster,omitempty"` // The URL of the poster image if requested. This will only be available if status is done.
	Rendertime float64 `json:"renderTime,omitempty"` // The time taken to render the asset in milliseconds.
	Created string `json:"created"` // The time the render task was initially queued.
	Status string `json:"status"` // The status of the render task. <ul> <li>`queued` - render is queued waiting to be rendered</li> <li>`fetching` - assets are being fetched</li> <li>`rendering` - the asset is being rendered</li> <li>`saving` - the final asset is being saved to storage</li> <li>`done` - the asset is ready to be downloaded</li> <li>`failed` - there was an error rendering the asset</li> </ul>
	Thumbnail string `json:"thumbnail,omitempty"` // The URL of the thumbnail image if requested. This will only be available if status is done.
	Duration float64 `json:"duration,omitempty"` // The output video or audio length in seconds.
	Owner string `json:"owner"` // The owner id of the render task.
	Plan string `json:"plan,omitempty"` // The customer subscription plan.
	Updated string `json:"updated"` // The time the render status was last updated.
}

// ImageAsset represents the ImageAsset schema from the OpenAPI specification
type ImageAsset struct {
	Crop Crop `json:"crop,omitempty"` // Crop the sides of an asset by a relative amount. The size of the crop is specified using a scale between 0 and 1, relative to the screen width - i.e a left crop of 0.5 will crop half of the asset from the left, a top crop of 0.25 will crop the top by quarter of the asset.
	Src string `json:"src"` // The image source URL. The URL must be publicly accessible or include credentials.
	TypeField string `json:"type"` // The type of asset - set to `image` for images.
}

// RenderResponse represents the RenderResponse schema from the OpenAPI specification
type RenderResponse struct {
	Response RenderResponseData `json:"response"` // The response data returned with the [RenderResponse](#tocs_renderresponse) including status and URL.
	Success bool `json:"success"` // `true` if status available, else `false`.
	Message string `json:"message"` // `OK` or an error message.
}

// Timeline represents the Timeline schema from the OpenAPI specification
type Timeline struct {
	Tracks []Track `json:"tracks"` // A timeline consists of an array of tracks, each track containing clips. Tracks are layered on top of each other in the same order they are added to the array with the top most track layered over the top of those below it. Ensure that a track containing titles is the top most track so that it is displayed above videos and images.
	Background string `json:"background,omitempty"` // A hexadecimal value for the timeline background colour. Defaults to #000000 (black).
	Cache bool `json:"cache,omitempty"` // Disable the caching of ingested source footage and assets. See [caching](https://shotstack.gitbook.io/docs/guides/architecting-an-application/caching) for more details.
	Fonts []Font `json:"fonts,omitempty"` // An array of custom fonts to be downloaded for use by the HTML assets.
	Soundtrack Soundtrack `json:"soundtrack,omitempty"` // A music or audio file in mp3 format that plays for the duration of the rendered video or the length of the audio file, which ever is shortest.
}
