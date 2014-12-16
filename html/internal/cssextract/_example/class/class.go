package class

import (
	"gopkg.in/go-on/lib.v2/html"
)

var (
	ContentLeft              = html.Class("content-left")
	ContentCenter            = html.Class("content-center")
	ContentRight             = html.Class("content-right")
	ShowAll                  = html.Class("show-all")
	HideAll                  = html.Class("hide-all")
	InkGrid                  = html.Class("ink-grid")
	ColumnGroup              = html.Class("column-group")
	Gutters                  = html.Class("gutters")
	HalfGutters              = html.Class("half-gutters")
	QuarterGutters           = html.Class("quarter-gutters")
	HorizontalGutters        = html.Class("horizontal-gutters")
	HalfHorizontalGutters    = html.Class("half-horizontal-gutters")
	QuarterHorizontalGutters = html.Class("quarter-horizontal-gutters")
	Large95                  = html.Class("large-95")
	Large90                  = html.Class("large-90")
	Large85                  = html.Class("large-85")
	Large80                  = html.Class("large-80")
	Large75                  = html.Class("large-75")
	Large70                  = html.Class("large-70")
	Large65                  = html.Class("large-65")
	Large60                  = html.Class("large-60")
	Large55                  = html.Class("large-55")
	Large50                  = html.Class("large-50")
	Large45                  = html.Class("large-45")
	Large40                  = html.Class("large-40")
	Large35                  = html.Class("large-35")
	Large30                  = html.Class("large-30")
	Large25                  = html.Class("large-25")
	Large20                  = html.Class("large-20")
	Large15                  = html.Class("large-15")
	Large10                  = html.Class("large-10")
	Large5                   = html.Class("large-5")
	Large33                  = html.Class("large-33")
	Large66                  = html.Class("large-66")
	Large100                 = html.Class("large-100")
	Animated                 = html.Class("animated")
	VerticalGutters          = html.Class("vertical-gutters")
	HalfVerticalGutters      = html.Class("half-vertical-gutters")
	LargePushLeft            = html.Class("large-push-left")
	LargePushCenter          = html.Class("large-push-center")
	LargePushRight           = html.Class("large-push-right")
	LargeAlignLeft           = html.Class("large-align-left")
	LargeAlignCenter         = html.Class("large-align-center")
	LargeAlignRight          = html.Class("large-align-right")
	Space                    = html.Class("space")
	HalfSpace                = html.Class("half-space")
	QuarterSpace             = html.Class("quarter-space")
	Vspace                   = html.Class("vspace")
	Hspace                   = html.Class("hspace")
	VerticalSpace            = html.Class("vertical-space")
	HalfVerticalSpace        = html.Class("half-vertical-space")
	QuarterVerticalSpace     = html.Class("quarter-vertical-space")
	HorizontalSpace          = html.Class("horizontal-space")
	HalfHorizontalSpace      = html.Class("half-horizontal-space")
	QuarterHorizontalSpace   = html.Class("quarter-horizontal-space")
	TopSpace                 = html.Class("top-space")
	HalfTopSpace             = html.Class("half-top-space")
	QuarterTopSpace          = html.Class("quarter-top-space")
	RightSpace               = html.Class("right-space")
	HalfRightSpace           = html.Class("half-right-space")
	QuarterRightSpace        = html.Class("quarter-right-space")
	BottomSpace              = html.Class("bottom-space")
	HalfBottomSpace          = html.Class("half-bottom-space")
	QuarterBottomSpace       = html.Class("quarter-bottom-space")
	LeftSpace                = html.Class("left-space")
	HalfLeftSpace            = html.Class("half-left-space")
	QuarterLeftSpace         = html.Class("quarter-left-space")
	HideLarge                = html.Class("hide-large")
	ShowLarge                = html.Class("show-large")
	Medium95                 = html.Class("medium-95")
	Medium90                 = html.Class("medium-90")
	Medium85                 = html.Class("medium-85")
	Medium80                 = html.Class("medium-80")
	Medium75                 = html.Class("medium-75")
	Medium70                 = html.Class("medium-70")
	Medium65                 = html.Class("medium-65")
	Medium60                 = html.Class("medium-60")
	Medium55                 = html.Class("medium-55")
	Medium50                 = html.Class("medium-50")
	Medium45                 = html.Class("medium-45")
	Medium40                 = html.Class("medium-40")
	Medium35                 = html.Class("medium-35")
	Medium30                 = html.Class("medium-30")
	Medium25                 = html.Class("medium-25")
	Medium20                 = html.Class("medium-20")
	Medium15                 = html.Class("medium-15")
	Medium10                 = html.Class("medium-10")
	Medium5                  = html.Class("medium-5")
	Medium33                 = html.Class("medium-33")
	Medium66                 = html.Class("medium-66")
	Medium100                = html.Class("medium-100")
	MediumPushLeft           = html.Class("medium-push-left")
	MediumPushCenter         = html.Class("medium-push-center")
	MediumPushRight          = html.Class("medium-push-right")
	MediumAlignLeft          = html.Class("medium-align-left")
	MediumAlignCenter        = html.Class("medium-align-center")
	MediumAlignRight         = html.Class("medium-align-right")
	HideMedium               = html.Class("hide-medium")
	ShowMedium               = html.Class("show-medium")
	Small95                  = html.Class("small-95")
	Small90                  = html.Class("small-90")
	Small85                  = html.Class("small-85")
	Small80                  = html.Class("small-80")
	Small75                  = html.Class("small-75")
	Small70                  = html.Class("small-70")
	Small65                  = html.Class("small-65")
	Small60                  = html.Class("small-60")
	Small55                  = html.Class("small-55")
	Small50                  = html.Class("small-50")
	Small45                  = html.Class("small-45")
	Small40                  = html.Class("small-40")
	Small35                  = html.Class("small-35")
	Small30                  = html.Class("small-30")
	Small25                  = html.Class("small-25")
	Small20                  = html.Class("small-20")
	Small15                  = html.Class("small-15")
	Small10                  = html.Class("small-10")
	Small5                   = html.Class("small-5")
	Small33                  = html.Class("small-33")
	Small66                  = html.Class("small-66")
	Small100                 = html.Class("small-100")
	SmallPushLeft            = html.Class("small-push-left")
	SmallPushCenter          = html.Class("small-push-center")
	SmallPushRight           = html.Class("small-push-right")
	SmallAlignLeft           = html.Class("small-align-left")
	SmallAlignCenter         = html.Class("small-align-center")
	SmallAlignRight          = html.Class("small-align-right")
	HideSmall                = html.Class("hide-small")
	ShowSmall                = html.Class("show-small")
	Sans                     = html.Class("sans")
	Serif                    = html.Class("serif")
	Unstyled                 = html.Class("unstyled")
	Inline                   = html.Class("inline")
	Note                     = html.Class("note")
	Small                    = html.Class("small")
	Medium                   = html.Class("medium")
	Large                    = html.Class("large")
	Extralarge               = html.Class("extralarge")
	Lead                     = html.Class("lead")
	InkLabel                 = html.Class("ink-label")
	Success                  = html.Class("success")
	Invert                   = html.Class("invert")
	Warning                  = html.Class("warning")
	Error                    = html.Class("error")
	Info                     = html.Class("info")
	IconLarge                = html.Class("icon-large")
	IconFixedWidth           = html.Class("icon-fixed-width")
	IconsUl                  = html.Class("icons-ul")
	IconLi                   = html.Class("icon-li")
	Hide                     = html.Class("hide")
	IconMuted                = html.Class("icon-muted")
	IconLight                = html.Class("icon-light")
	IconDark                 = html.Class("icon-dark")
	IconBorder               = html.Class("icon-border")
	Icon2x                   = html.Class("icon-2x")
	Icon3x                   = html.Class("icon-3x")
	Icon4x                   = html.Class("icon-4x")
	Icon5x                   = html.Class("icon-5x")
	PullRight                = html.Class("pull-right")
	PullLeft                 = html.Class("pull-left")
	IconStack                = html.Class("icon-stack")
	IconStackBase            = html.Class("icon-stack-base")
	IconSpin                 = html.Class("icon-spin")
	IconRotate90             = html.Class("icon-rotate-90")
	IconRotate180            = html.Class("icon-rotate-180")
	IconRotate270            = html.Class("icon-rotate-270")
	IconFlipHorizontal       = html.Class("icon-flip-horizontal")
	IconFlipVertical         = html.Class("icon-flip-vertical")
	IconGlass                = html.Class("icon-glass")
	IconMusic                = html.Class("icon-music")
	IconSearch               = html.Class("icon-search")
	IconEnvelopeAlt          = html.Class("icon-envelope-alt")
	IconHeart                = html.Class("icon-heart")
	IconStar                 = html.Class("icon-star")
	IconStarEmpty            = html.Class("icon-star-empty")
	IconUser                 = html.Class("icon-user")
	IconFilm                 = html.Class("icon-film")
	IconThLarge              = html.Class("icon-th-large")
	IconTh                   = html.Class("icon-th")
	IconThList               = html.Class("icon-th-list")
	IconOk                   = html.Class("icon-ok")
	IconRemove               = html.Class("icon-remove")
	IconZoomIn               = html.Class("icon-zoom-in")
	IconZoomOut              = html.Class("icon-zoom-out")
	IconPowerOff             = html.Class("icon-power-off")
	IconOff                  = html.Class("icon-off")
	IconSignal               = html.Class("icon-signal")
	IconGear                 = html.Class("icon-gear")
	IconCog                  = html.Class("icon-cog")
	IconTrash                = html.Class("icon-trash")
	IconHome                 = html.Class("icon-home")
	IconFileAlt              = html.Class("icon-file-alt")
	IconTime                 = html.Class("icon-time")
	IconRoad                 = html.Class("icon-road")
	IconDownloadAlt          = html.Class("icon-download-alt")
	IconDownload             = html.Class("icon-download")
	IconUpload               = html.Class("icon-upload")
	IconInbox                = html.Class("icon-inbox")
	IconPlayCircle           = html.Class("icon-play-circle")
	IconRotateRight          = html.Class("icon-rotate-right")
	IconRepeat               = html.Class("icon-repeat")
	IconRefresh              = html.Class("icon-refresh")
	IconListAlt              = html.Class("icon-list-alt")
	IconLock                 = html.Class("icon-lock")
	IconFlag                 = html.Class("icon-flag")
	IconHeadphones           = html.Class("icon-headphones")
	IconVolumeOff            = html.Class("icon-volume-off")
	IconVolumeDown           = html.Class("icon-volume-down")
	IconVolumeUp             = html.Class("icon-volume-up")
	IconQrcode               = html.Class("icon-qrcode")
	IconBarcode              = html.Class("icon-barcode")
	IconTag                  = html.Class("icon-tag")
	IconTags                 = html.Class("icon-tags")
	IconBook                 = html.Class("icon-book")
	IconBookmark             = html.Class("icon-bookmark")
	IconPrint                = html.Class("icon-print")
	IconCamera               = html.Class("icon-camera")
	IconFont                 = html.Class("icon-font")
	IconBold                 = html.Class("icon-bold")
	IconItalic               = html.Class("icon-italic")
	IconTextHeight           = html.Class("icon-text-height")
	IconTextWidth            = html.Class("icon-text-width")
	IconAlignLeft            = html.Class("icon-align-left")
	IconAlignCenter          = html.Class("icon-align-center")
	IconAlignRight           = html.Class("icon-align-right")
	IconAlignJustify         = html.Class("icon-align-justify")
	IconList                 = html.Class("icon-list")
	IconIndentLeft           = html.Class("icon-indent-left")
	IconIndentRight          = html.Class("icon-indent-right")
	IconFacetimeVideo        = html.Class("icon-facetime-video")
	IconPicture              = html.Class("icon-picture")
	IconPencil               = html.Class("icon-pencil")
	IconMapMarker            = html.Class("icon-map-marker")
	IconAdjust               = html.Class("icon-adjust")
	IconTint                 = html.Class("icon-tint")
	IconEdit                 = html.Class("icon-edit")
	IconShare                = html.Class("icon-share")
	IconCheck                = html.Class("icon-check")
	IconMove                 = html.Class("icon-move")
	IconStepBackward         = html.Class("icon-step-backward")
	IconFastBackward         = html.Class("icon-fast-backward")
	IconBackward             = html.Class("icon-backward")
	IconPlay                 = html.Class("icon-play")
	IconPause                = html.Class("icon-pause")
	IconStop                 = html.Class("icon-stop")
	IconForward              = html.Class("icon-forward")
	IconFastForward          = html.Class("icon-fast-forward")
	IconStepForward          = html.Class("icon-step-forward")
	IconEject                = html.Class("icon-eject")
	IconChevronLeft          = html.Class("icon-chevron-left")
	IconChevronRight         = html.Class("icon-chevron-right")
	IconPlusSign             = html.Class("icon-plus-sign")
	IconMinusSign            = html.Class("icon-minus-sign")
	IconRemoveSign           = html.Class("icon-remove-sign")
	IconOkSign               = html.Class("icon-ok-sign")
	IconQuestionSign         = html.Class("icon-question-sign")
	IconInfoSign             = html.Class("icon-info-sign")
	IconScreenshot           = html.Class("icon-screenshot")
	IconRemoveCircle         = html.Class("icon-remove-circle")
	IconOkCircle             = html.Class("icon-ok-circle")
	IconBanCircle            = html.Class("icon-ban-circle")
	IconArrowLeft            = html.Class("icon-arrow-left")
	IconArrowRight           = html.Class("icon-arrow-right")
	IconArrowUp              = html.Class("icon-arrow-up")
	IconArrowDown            = html.Class("icon-arrow-down")
	IconMailForward          = html.Class("icon-mail-forward")
	IconShareAlt             = html.Class("icon-share-alt")
	IconResizeFull           = html.Class("icon-resize-full")
	IconResizeSmall          = html.Class("icon-resize-small")
	IconPlus                 = html.Class("icon-plus")
	IconMinus                = html.Class("icon-minus")
	IconAsterisk             = html.Class("icon-asterisk")
	IconExclamationSign      = html.Class("icon-exclamation-sign")
	IconGift                 = html.Class("icon-gift")
	IconLeaf                 = html.Class("icon-leaf")
	IconFire                 = html.Class("icon-fire")
	IconEyeOpen              = html.Class("icon-eye-open")
	IconEyeClose             = html.Class("icon-eye-close")
	IconWarningSign          = html.Class("icon-warning-sign")
	IconPlane                = html.Class("icon-plane")
	IconCalendar             = html.Class("icon-calendar")
	IconRandom               = html.Class("icon-random")
	IconComment              = html.Class("icon-comment")
	IconMagnet               = html.Class("icon-magnet")
	IconChevronUp            = html.Class("icon-chevron-up")
	IconChevronDown          = html.Class("icon-chevron-down")
	IconRetweet              = html.Class("icon-retweet")
	IconShoppingCart         = html.Class("icon-shopping-cart")
	IconFolderClose          = html.Class("icon-folder-close")
	IconFolderOpen           = html.Class("icon-folder-open")
	IconResizeVertical       = html.Class("icon-resize-vertical")
	IconResizeHorizontal     = html.Class("icon-resize-horizontal")
	IconBarChart             = html.Class("icon-bar-chart")
	IconTwitterSign          = html.Class("icon-twitter-sign")
	IconFacebookSign         = html.Class("icon-facebook-sign")
	IconCameraRetro          = html.Class("icon-camera-retro")
	IconKey                  = html.Class("icon-key")
	IconGears                = html.Class("icon-gears")
	IconCogs                 = html.Class("icon-cogs")
	IconComments             = html.Class("icon-comments")
	IconThumbsUpAlt          = html.Class("icon-thumbs-up-alt")
	IconThumbsDownAlt        = html.Class("icon-thumbs-down-alt")
	IconStarHalf             = html.Class("icon-star-half")
	IconHeartEmpty           = html.Class("icon-heart-empty")
	IconSignout              = html.Class("icon-signout")
	IconLinkedinSign         = html.Class("icon-linkedin-sign")
	IconPushpin              = html.Class("icon-pushpin")
	IconExternalLink         = html.Class("icon-external-link")
	IconSignin               = html.Class("icon-signin")
	IconTrophy               = html.Class("icon-trophy")
	IconGithubSign           = html.Class("icon-github-sign")
	IconUploadAlt            = html.Class("icon-upload-alt")
	IconLemon                = html.Class("icon-lemon")
	IconPhone                = html.Class("icon-phone")
	IconUnchecked            = html.Class("icon-unchecked")
	IconCheckEmpty           = html.Class("icon-check-empty")
	IconBookmarkEmpty        = html.Class("icon-bookmark-empty")
	IconPhoneSign            = html.Class("icon-phone-sign")
	IconTwitter              = html.Class("icon-twitter")
	IconFacebook             = html.Class("icon-facebook")
	IconGithub               = html.Class("icon-github")
	IconUnlock               = html.Class("icon-unlock")
	IconCreditCard           = html.Class("icon-credit-card")
	IconRss                  = html.Class("icon-rss")
	IconHdd                  = html.Class("icon-hdd")
	IconBullhorn             = html.Class("icon-bullhorn")
	IconBell                 = html.Class("icon-bell")
	IconCertificate          = html.Class("icon-certificate")
	IconHandRight            = html.Class("icon-hand-right")
	IconHandLeft             = html.Class("icon-hand-left")
	IconHandUp               = html.Class("icon-hand-up")
	IconHandDown             = html.Class("icon-hand-down")
	IconCircleArrowLeft      = html.Class("icon-circle-arrow-left")
	IconCircleArrowRight     = html.Class("icon-circle-arrow-right")
	IconCircleArrowUp        = html.Class("icon-circle-arrow-up")
	IconCircleArrowDown      = html.Class("icon-circle-arrow-down")
	IconGlobe                = html.Class("icon-globe")
	IconWrench               = html.Class("icon-wrench")
	IconTasks                = html.Class("icon-tasks")
	IconFilter               = html.Class("icon-filter")
	IconBriefcase            = html.Class("icon-briefcase")
	IconFullscreen           = html.Class("icon-fullscreen")
	IconGroup                = html.Class("icon-group")
	IconLink                 = html.Class("icon-link")
	IconCloud                = html.Class("icon-cloud")
	IconBeaker               = html.Class("icon-beaker")
	IconCut                  = html.Class("icon-cut")
	IconCopy                 = html.Class("icon-copy")
	IconPaperclip            = html.Class("icon-paperclip")
	IconPaperClip            = html.Class("icon-paper-clip")
	IconSave                 = html.Class("icon-save")
	IconSignBlank            = html.Class("icon-sign-blank")
	IconReorder              = html.Class("icon-reorder")
	IconListUl               = html.Class("icon-list-ul")
	IconListOl               = html.Class("icon-list-ol")
	IconStrikethrough        = html.Class("icon-strikethrough")
	IconUnderline            = html.Class("icon-underline")
	IconTable                = html.Class("icon-table")
	IconMagic                = html.Class("icon-magic")
	IconTruck                = html.Class("icon-truck")
	IconPinterest            = html.Class("icon-pinterest")
	IconPinterestSign        = html.Class("icon-pinterest-sign")
	IconGooglePlusSign       = html.Class("icon-google-plus-sign")
	IconGooglePlus           = html.Class("icon-google-plus")
	IconMoney                = html.Class("icon-money")
	IconCaretDown            = html.Class("icon-caret-down")
	IconCaretUp              = html.Class("icon-caret-up")
	IconCaretLeft            = html.Class("icon-caret-left")
	IconCaretRight           = html.Class("icon-caret-right")
	IconColumns              = html.Class("icon-columns")
	IconSort                 = html.Class("icon-sort")
	IconSortDown             = html.Class("icon-sort-down")
	IconSortUp               = html.Class("icon-sort-up")
	IconEnvelope             = html.Class("icon-envelope")
	IconLinkedin             = html.Class("icon-linkedin")
	IconRotateLeft           = html.Class("icon-rotate-left")
	IconUndo                 = html.Class("icon-undo")
	IconLegal                = html.Class("icon-legal")
	IconDashboard            = html.Class("icon-dashboard")
	IconCommentAlt           = html.Class("icon-comment-alt")
	IconCommentsAlt          = html.Class("icon-comments-alt")
	IconBolt                 = html.Class("icon-bolt")
	IconSitemap              = html.Class("icon-sitemap")
	IconUmbrella             = html.Class("icon-umbrella")
	IconPaste                = html.Class("icon-paste")
	IconLightbulb            = html.Class("icon-lightbulb")
	IconExchange             = html.Class("icon-exchange")
	IconCloudDownload        = html.Class("icon-cloud-download")
	IconCloudUpload          = html.Class("icon-cloud-upload")
	IconUserMd               = html.Class("icon-user-md")
	IconStethoscope          = html.Class("icon-stethoscope")
	IconSuitcase             = html.Class("icon-suitcase")
	IconBellAlt              = html.Class("icon-bell-alt")
	IconCoffee               = html.Class("icon-coffee")
	IconFood                 = html.Class("icon-food")
	IconFileTextAlt          = html.Class("icon-file-text-alt")
	IconBuilding             = html.Class("icon-building")
	IconHospital             = html.Class("icon-hospital")
	IconAmbulance            = html.Class("icon-ambulance")
	IconMedkit               = html.Class("icon-medkit")
	IconFighterJet           = html.Class("icon-fighter-jet")
	IconBeer                 = html.Class("icon-beer")
	IconHSign                = html.Class("icon-h-sign")
	IconPlusSignAlt          = html.Class("icon-plus-sign-alt")
	IconDoubleAngleLeft      = html.Class("icon-double-angle-left")
	IconDoubleAngleRight     = html.Class("icon-double-angle-right")
	IconDoubleAngleUp        = html.Class("icon-double-angle-up")
	IconDoubleAngleDown      = html.Class("icon-double-angle-down")
	IconAngleLeft            = html.Class("icon-angle-left")
	IconAngleRight           = html.Class("icon-angle-right")
	IconAngleUp              = html.Class("icon-angle-up")
	IconAngleDown            = html.Class("icon-angle-down")
	IconDesktop              = html.Class("icon-desktop")
	IconLaptop               = html.Class("icon-laptop")
	IconTablet               = html.Class("icon-tablet")
	IconMobilePhone          = html.Class("icon-mobile-phone")
	IconCircleBlank          = html.Class("icon-circle-blank")
	IconQuoteLeft            = html.Class("icon-quote-left")
	IconQuoteRight           = html.Class("icon-quote-right")
	IconSpinner              = html.Class("icon-spinner")
	IconCircle               = html.Class("icon-circle")
	IconMailReply            = html.Class("icon-mail-reply")
	IconReply                = html.Class("icon-reply")
	IconGithubAlt            = html.Class("icon-github-alt")
	IconFolderCloseAlt       = html.Class("icon-folder-close-alt")
	IconFolderOpenAlt        = html.Class("icon-folder-open-alt")
	IconExpandAlt            = html.Class("icon-expand-alt")
	IconCollapseAlt          = html.Class("icon-collapse-alt")
	IconSmile                = html.Class("icon-smile")
	IconFrown                = html.Class("icon-frown")
	IconMeh                  = html.Class("icon-meh")
	IconGamepad              = html.Class("icon-gamepad")
	IconKeyboard             = html.Class("icon-keyboard")
	IconFlagAlt              = html.Class("icon-flag-alt")
	IconFlagCheckered        = html.Class("icon-flag-checkered")
	IconTerminal             = html.Class("icon-terminal")
	IconCode                 = html.Class("icon-code")
	IconReplyAll             = html.Class("icon-reply-all")
	IconMailReplyAll         = html.Class("icon-mail-reply-all")
	IconStarHalfFull         = html.Class("icon-star-half-full")
	IconStarHalfEmpty        = html.Class("icon-star-half-empty")
	IconLocationArrow        = html.Class("icon-location-arrow")
	IconCrop                 = html.Class("icon-crop")
	IconCodeFork             = html.Class("icon-code-fork")
	IconUnlink               = html.Class("icon-unlink")
	IconQuestion             = html.Class("icon-question")
	IconInfo                 = html.Class("icon-info")
	IconExclamation          = html.Class("icon-exclamation")
	IconSuperscript          = html.Class("icon-superscript")
	IconSubscript            = html.Class("icon-subscript")
	IconEraser               = html.Class("icon-eraser")
	IconPuzzlePiece          = html.Class("icon-puzzle-piece")
	IconMicrophone           = html.Class("icon-microphone")
	IconMicrophoneOff        = html.Class("icon-microphone-off")
	IconShield               = html.Class("icon-shield")
	IconCalendarEmpty        = html.Class("icon-calendar-empty")
	IconFireExtinguisher     = html.Class("icon-fire-extinguisher")
	IconRocket               = html.Class("icon-rocket")
	IconMaxcdn               = html.Class("icon-maxcdn")
	IconChevronSignLeft      = html.Class("icon-chevron-sign-left")
	IconChevronSignRight     = html.Class("icon-chevron-sign-right")
	IconChevronSignUp        = html.Class("icon-chevron-sign-up")
	IconChevronSignDown      = html.Class("icon-chevron-sign-down")
	IconHtml5                = html.Class("icon-html5")
	IconCss3                 = html.Class("icon-css3")
	IconAnchor               = html.Class("icon-anchor")
	IconUnlockAlt            = html.Class("icon-unlock-alt")
	IconBullseye             = html.Class("icon-bullseye")
	IconEllipsisHorizontal   = html.Class("icon-ellipsis-horizontal")
	IconEllipsisVertical     = html.Class("icon-ellipsis-vertical")
	IconRssSign              = html.Class("icon-rss-sign")
	IconPlaySign             = html.Class("icon-play-sign")
	IconTicket               = html.Class("icon-ticket")
	IconMinusSignAlt         = html.Class("icon-minus-sign-alt")
	IconCheckMinus           = html.Class("icon-check-minus")
	IconLevelUp              = html.Class("icon-level-up")
	IconLevelDown            = html.Class("icon-level-down")
	IconCheckSign            = html.Class("icon-check-sign")
	IconEditSign             = html.Class("icon-edit-sign")
	IconExternalLinkSign     = html.Class("icon-external-link-sign")
	IconShareSign            = html.Class("icon-share-sign")
	IconCompass              = html.Class("icon-compass")
	IconCollapse             = html.Class("icon-collapse")
	IconCollapseTop          = html.Class("icon-collapse-top")
	IconExpand               = html.Class("icon-expand")
	IconEuro                 = html.Class("icon-euro")
	IconEur                  = html.Class("icon-eur")
	IconGbp                  = html.Class("icon-gbp")
	IconDollar               = html.Class("icon-dollar")
	IconUsd                  = html.Class("icon-usd")
	IconRupee                = html.Class("icon-rupee")
	IconInr                  = html.Class("icon-inr")
	IconYen                  = html.Class("icon-yen")
	IconJpy                  = html.Class("icon-jpy")
	IconRenminbi             = html.Class("icon-renminbi")
	IconCny                  = html.Class("icon-cny")
	IconWon                  = html.Class("icon-won")
	IconKrw                  = html.Class("icon-krw")
	IconBitcoin              = html.Class("icon-bitcoin")
	IconBtc                  = html.Class("icon-btc")
	IconFile                 = html.Class("icon-file")
	IconFileText             = html.Class("icon-file-text")
	IconSortByAlphabet       = html.Class("icon-sort-by-alphabet")
	IconSortByAlphabetAlt    = html.Class("icon-sort-by-alphabet-alt")
	IconSortByAttributes     = html.Class("icon-sort-by-attributes")
	IconSortByAttributesAlt  = html.Class("icon-sort-by-attributes-alt")
	IconSortByOrder          = html.Class("icon-sort-by-order")
	IconSortByOrderAlt       = html.Class("icon-sort-by-order-alt")
	IconThumbsUp             = html.Class("icon-thumbs-up")
	IconThumbsDown           = html.Class("icon-thumbs-down")
	IconYoutubeSign          = html.Class("icon-youtube-sign")
	IconYoutube              = html.Class("icon-youtube")
	IconXing                 = html.Class("icon-xing")
	IconXingSign             = html.Class("icon-xing-sign")
	IconYoutubePlay          = html.Class("icon-youtube-play")
	IconDropbox              = html.Class("icon-dropbox")
	IconStackexchange        = html.Class("icon-stackexchange")
	IconInstagram            = html.Class("icon-instagram")
	IconFlickr               = html.Class("icon-flickr")
	IconAdn                  = html.Class("icon-adn")
	IconBitbucket            = html.Class("icon-bitbucket")
	IconBitbucketSign        = html.Class("icon-bitbucket-sign")
	IconTumblr               = html.Class("icon-tumblr")
	IconTumblrSign           = html.Class("icon-tumblr-sign")
	IconLongArrowDown        = html.Class("icon-long-arrow-down")
	IconLongArrowUp          = html.Class("icon-long-arrow-up")
	IconLongArrowLeft        = html.Class("icon-long-arrow-left")
	IconLongArrowRight       = html.Class("icon-long-arrow-right")
	IconApple                = html.Class("icon-apple")
	IconWindows              = html.Class("icon-windows")
	IconAndroid              = html.Class("icon-android")
	IconLinux                = html.Class("icon-linux")
	IconDribbble             = html.Class("icon-dribbble")
	IconSkype                = html.Class("icon-skype")
	IconFoursquare           = html.Class("icon-foursquare")
	IconTrello               = html.Class("icon-trello")
	IconFemale               = html.Class("icon-female")
	IconMale                 = html.Class("icon-male")
	IconGittip               = html.Class("icon-gittip")
	IconSun                  = html.Class("icon-sun")
	IconMoon                 = html.Class("icon-moon")
	IconArchive              = html.Class("icon-archive")
	IconBug                  = html.Class("icon-bug")
	IconVk                   = html.Class("icon-vk")
	IconWeibo                = html.Class("icon-weibo")
	IconRenren               = html.Class("icon-renren")
	InkNavigation            = html.Class("ink-navigation")
	Menu                     = html.Class("menu")
	Submenu                  = html.Class("submenu")
	Horizontal               = html.Class("horizontal")
	Control                  = html.Class("control")
	Vertical                 = html.Class("vertical")
	Dropdown                 = html.Class("dropdown")
	Breadcrumbs              = html.Class("breadcrumbs")
	Active                   = html.Class("active")
	Pagination               = html.Class("pagination")
	Disabled                 = html.Class("disabled")
	Pills                    = html.Class("pills")
	InkDropdown              = html.Class("ink-dropdown")
	DropdownMenu             = html.Class("dropdown-menu")
	SeparatorAbove           = html.Class("separator-above")
	SeparatorBelow           = html.Class("separator-below")
	Heading                  = html.Class("heading")
	White                    = html.Class("white")
	Grey                     = html.Class("grey")
	Black                    = html.Class("black")
	Orange                   = html.Class("orange")
	Blue                     = html.Class("blue")
	Green                    = html.Class("green")
	Red                      = html.Class("red")
	Flat                     = html.Class("flat")
	Rounded                  = html.Class("rounded")
	Shadowed                 = html.Class("shadowed")
	InkForm                  = html.Class("ink-form")
	Tip                      = html.Class("tip")
	Label                    = html.Class("label")
	InputFile                = html.Class("input-file")
	InkButton                = html.Class("ink-button")
	ControlGroup             = html.Class("control-group")
	AppendButton             = html.Class("append-button")
	AppendSymbol             = html.Class("append-symbol")
	PrependButton            = html.Class("prepend-button")
	PrependSymbol            = html.Class("prepend-symbol")
	Validation               = html.Class("validation")
	Required                 = html.Class("required")
	StatusIndicator          = html.Class("status-indicator")
	InkAlert                 = html.Class("ink-alert")
	Basic                    = html.Class("basic")
	Block                    = html.Class("block")
	InkClose                 = html.Class("ink-close")
	InkDismiss               = html.Class("ink-dismiss")
	InkBadge                 = html.Class("ink-badge")
	InkTooltip               = html.Class("ink-tooltip")
	Content                  = html.Class("content")
	Arrow                    = html.Class("arrow")
	Up                       = html.Class("up")
	Down                     = html.Class("down")
	Left                     = html.Class("left")
	Right                    = html.Class("right")
	InkDisabled              = html.Class("ink-disabled")
	ButtonGroup              = html.Class("button-group")
	ButtonToolbar            = html.Class("button-toolbar")
	InkTable                 = html.Class("ink-table")
	Alternating              = html.Class("alternating")
	Hover                    = html.Class("hover")
	Bordered                 = html.Class("bordered")
	InkGallery               = html.Class("ink-gallery")
	Thumbs                   = html.Class("thumbs")
	Slider                   = html.Class("slider")
	ArticleText              = html.Class("article_text")
	Example1                 = html.Class("example1")
	Example2                 = html.Class("example2")
	Stage                    = html.Class("stage")
	Next                     = html.Class("next")
	Previous                 = html.Class("previous")
	RightNav                 = html.Class("rightNav")
	SapoComponentDatepicker  = html.Class("sapo_component_datepicker")
	SapoCalTopOptions        = html.Class("sapo_cal_top_options")
	Clean                    = html.Class("clean")
	Close                    = html.Class("close")
	SapoCalTop               = html.Class("sapo_cal_top")
	SapoCalPrev              = html.Class("sapo_cal_prev")
	SapoCalNext              = html.Class("sapo_cal_next")
	SapoCalMonthDesc         = html.Class("sapo_cal_month_desc")
	SapoCalMonth             = html.Class("sapo_cal_month")
	SapoCalYearSelector      = html.Class("sapo_cal_year_selector")
	SapoCalMonthSelector     = html.Class("sapo_cal_month_selector")
	SapoCalOn                = html.Class("sapo_cal_on")
	SapoCalOff               = html.Class("sapo_cal_off")
	SapoCalHeader            = html.Class("sapo_cal_header")
	SapoCalMiddle            = html.Class("sapo_cal_middle")
	InkModal                 = html.Class("ink-modal")
	ModalBody                = html.Class("modal-body")
	ModalHeader              = html.Class("modal-header")
	ModalClose               = html.Class("modal-close")
	ModalFooter              = html.Class("modal-footer")
	InkModalOpen             = html.Class("ink-modal-open")
	InkProgressBar           = html.Class("ink-progress-bar")
	Caption                  = html.Class("caption")
	Bar                      = html.Class("bar")
	InkTabs                  = html.Class("ink-tabs")
	TabsNav                  = html.Class("tabs-nav")
	TabsContent              = html.Class("tabs-content")
	Top                      = html.Class("top")
	Bottom                   = html.Class("bottom")
	InkSortableList          = html.Class("ink-sortable-list")
	InkTreeView              = html.Class("ink-tree-view")
	Open                     = html.Class("open")
	Closed                   = html.Class("closed")
	InkCarousel              = html.Class("ink-carousel")
	Slide                    = html.Class("slide")
	Hider                    = html.Class("hider")
	CaptionOverTop           = html.Class("caption-over-top")
	CaptionOverBottom        = html.Class("caption-over-bottom")
	Light                    = html.Class("light")
	Dark                     = html.Class("dark")
	PushLeft                 = html.Class("push-left")
	PushRight                = html.Class("push-right")
	PushCenter               = html.Class("push-center")
	Clearfix                 = html.Class("clearfix")
	NoMargin                 = html.Class("no-margin")
	ScreenSizeHelper         = html.Class("screen-size-helper")
	Title                    = html.Class("title")
	Drag                     = html.Class("drag")
	InkShade                 = html.Class("ink-shade")
	Fade                     = html.Class("fade")
	Visible                  = html.Class("visible")
)
