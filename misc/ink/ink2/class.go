package ink2

import (
	html "github.com/go-on/lib/types"
)

var (
	Content_left                = html.Class("content-left")
	Content_center              = html.Class("content-center")
	Content_right               = html.Class("content-right")
	Show_all                    = html.Class("show-all")
	Hide_all                    = html.Class("hide-all")
	Ink_grid                    = html.Class("ink-grid")
	Column_group                = html.Class("column-group")
	Gutters                     = html.Class("gutters")
	Half_gutters                = html.Class("half-gutters")
	Quarter_gutters             = html.Class("quarter-gutters")
	Horizontal_gutters          = html.Class("horizontal-gutters")
	Half_horizontal_gutters     = html.Class("half-horizontal-gutters")
	Quarter_horizontal_gutters  = html.Class("quarter-horizontal-gutters")
	Large_95                    = html.Class("large-95")
	Large_90                    = html.Class("large-90")
	Large_85                    = html.Class("large-85")
	Large_80                    = html.Class("large-80")
	Large_75                    = html.Class("large-75")
	Large_70                    = html.Class("large-70")
	Large_65                    = html.Class("large-65")
	Large_60                    = html.Class("large-60")
	Large_55                    = html.Class("large-55")
	Large_50                    = html.Class("large-50")
	Large_45                    = html.Class("large-45")
	Large_40                    = html.Class("large-40")
	Large_35                    = html.Class("large-35")
	Large_30                    = html.Class("large-30")
	Large_25                    = html.Class("large-25")
	Large_20                    = html.Class("large-20")
	Large_15                    = html.Class("large-15")
	Large_10                    = html.Class("large-10")
	Large_5                     = html.Class("large-5")
	Large_33                    = html.Class("large-33")
	Large_66                    = html.Class("large-66")
	Large_100                   = html.Class("large-100")
	Animated                    = html.Class("animated")
	Vertical_gutters            = html.Class("vertical-gutters")
	Half_vertical_gutters       = html.Class("half-vertical-gutters")
	Large_push_left             = html.Class("large-push-left")
	Large_push_center           = html.Class("large-push-center")
	Large_push_right            = html.Class("large-push-right")
	Large_align_left            = html.Class("large-align-left")
	Large_align_center          = html.Class("large-align-center")
	Large_align_right           = html.Class("large-align-right")
	Space                       = html.Class("space")
	Half_space                  = html.Class("half-space")
	Quarter_space               = html.Class("quarter-space")
	Vspace                      = html.Class("vspace")
	Hspace                      = html.Class("hspace")
	Vertical_space              = html.Class("vertical-space")
	Half_vertical_space         = html.Class("half-vertical-space")
	Quarter_vertical_space      = html.Class("quarter-vertical-space")
	Horizontal_space            = html.Class("horizontal-space")
	Half_horizontal_space       = html.Class("half-horizontal-space")
	Quarter_horizontal_space    = html.Class("quarter-horizontal-space")
	Top_space                   = html.Class("top-space")
	Half_top_space              = html.Class("half-top-space")
	Quarter_top_space           = html.Class("quarter-top-space")
	Right_space                 = html.Class("right-space")
	Half_right_space            = html.Class("half-right-space")
	Quarter_right_space         = html.Class("quarter-right-space")
	Bottom_space                = html.Class("bottom-space")
	Half_bottom_space           = html.Class("half-bottom-space")
	Quarter_bottom_space        = html.Class("quarter-bottom-space")
	Left_space                  = html.Class("left-space")
	Half_left_space             = html.Class("half-left-space")
	Quarter_left_space          = html.Class("quarter-left-space")
	Hide_large                  = html.Class("hide-large")
	Show_large                  = html.Class("show-large")
	Medium_95                   = html.Class("medium-95")
	Medium_90                   = html.Class("medium-90")
	Medium_85                   = html.Class("medium-85")
	Medium_80                   = html.Class("medium-80")
	Medium_75                   = html.Class("medium-75")
	Medium_70                   = html.Class("medium-70")
	Medium_65                   = html.Class("medium-65")
	Medium_60                   = html.Class("medium-60")
	Medium_55                   = html.Class("medium-55")
	Medium_50                   = html.Class("medium-50")
	Medium_45                   = html.Class("medium-45")
	Medium_40                   = html.Class("medium-40")
	Medium_35                   = html.Class("medium-35")
	Medium_30                   = html.Class("medium-30")
	Medium_25                   = html.Class("medium-25")
	Medium_20                   = html.Class("medium-20")
	Medium_15                   = html.Class("medium-15")
	Medium_10                   = html.Class("medium-10")
	Medium_5                    = html.Class("medium-5")
	Medium_33                   = html.Class("medium-33")
	Medium_66                   = html.Class("medium-66")
	Medium_100                  = html.Class("medium-100")
	Medium_push_left            = html.Class("medium-push-left")
	Medium_push_center          = html.Class("medium-push-center")
	Medium_push_right           = html.Class("medium-push-right")
	Medium_align_left           = html.Class("medium-align-left")
	Medium_align_center         = html.Class("medium-align-center")
	Medium_align_right          = html.Class("medium-align-right")
	Hide_medium                 = html.Class("hide-medium")
	Show_medium                 = html.Class("show-medium")
	Small_95                    = html.Class("small-95")
	Small_90                    = html.Class("small-90")
	Small_85                    = html.Class("small-85")
	Small_80                    = html.Class("small-80")
	Small_75                    = html.Class("small-75")
	Small_70                    = html.Class("small-70")
	Small_65                    = html.Class("small-65")
	Small_60                    = html.Class("small-60")
	Small_55                    = html.Class("small-55")
	Small_50                    = html.Class("small-50")
	Small_45                    = html.Class("small-45")
	Small_40                    = html.Class("small-40")
	Small_35                    = html.Class("small-35")
	Small_30                    = html.Class("small-30")
	Small_25                    = html.Class("small-25")
	Small_20                    = html.Class("small-20")
	Small_15                    = html.Class("small-15")
	Small_10                    = html.Class("small-10")
	Small_5                     = html.Class("small-5")
	Small_33                    = html.Class("small-33")
	Small_66                    = html.Class("small-66")
	Small_100                   = html.Class("small-100")
	Small_push_left             = html.Class("small-push-left")
	Small_push_center           = html.Class("small-push-center")
	Small_push_right            = html.Class("small-push-right")
	Small_align_left            = html.Class("small-align-left")
	Small_align_center          = html.Class("small-align-center")
	Small_align_right           = html.Class("small-align-right")
	Hide_small                  = html.Class("hide-small")
	Show_small                  = html.Class("show-small")
	Sans                        = html.Class("sans")
	Serif                       = html.Class("serif")
	Unstyled                    = html.Class("unstyled")
	Inline                      = html.Class("inline")
	Note                        = html.Class("note")
	Small                       = html.Class("small")
	Medium                      = html.Class("medium")
	Large                       = html.Class("large")
	Extralarge                  = html.Class("extralarge")
	Lead                        = html.Class("lead")
	Ink_label                   = html.Class("ink-label")
	Success                     = html.Class("success")
	Invert                      = html.Class("invert")
	Warning                     = html.Class("warning")
	Error                       = html.Class("error")
	Info                        = html.Class("info")
	Icon_large                  = html.Class("icon-large")
	Icon_fixed_width            = html.Class("icon-fixed-width")
	Icons_ul                    = html.Class("icons-ul")
	Icon_li                     = html.Class("icon-li")
	Hide                        = html.Class("hide")
	Icon_muted                  = html.Class("icon-muted")
	Icon_light                  = html.Class("icon-light")
	Icon_dark                   = html.Class("icon-dark")
	Icon_border                 = html.Class("icon-border")
	Icon_2x                     = html.Class("icon-2x")
	Icon_3x                     = html.Class("icon-3x")
	Icon_4x                     = html.Class("icon-4x")
	Icon_5x                     = html.Class("icon-5x")
	Pull_right                  = html.Class("pull-right")
	Pull_left                   = html.Class("pull-left")
	Icon_stack                  = html.Class("icon-stack")
	Icon_stack_base             = html.Class("icon-stack-base")
	Icon_spin                   = html.Class("icon-spin")
	Icon_rotate_90              = html.Class("icon-rotate-90")
	Icon_rotate_180             = html.Class("icon-rotate-180")
	Icon_rotate_270             = html.Class("icon-rotate-270")
	Icon_flip_horizontal        = html.Class("icon-flip-horizontal")
	Icon_flip_vertical          = html.Class("icon-flip-vertical")
	Icon_glass                  = html.Class("icon-glass")
	Icon_music                  = html.Class("icon-music")
	Icon_search                 = html.Class("icon-search")
	Icon_envelope_alt           = html.Class("icon-envelope-alt")
	Icon_heart                  = html.Class("icon-heart")
	Icon_star                   = html.Class("icon-star")
	Icon_star_empty             = html.Class("icon-star-empty")
	Icon_user                   = html.Class("icon-user")
	Icon_film                   = html.Class("icon-film")
	Icon_th_large               = html.Class("icon-th-large")
	Icon_th                     = html.Class("icon-th")
	Icon_th_list                = html.Class("icon-th-list")
	Icon_ok                     = html.Class("icon-ok")
	Icon_remove                 = html.Class("icon-remove")
	Icon_zoom_in                = html.Class("icon-zoom-in")
	Icon_zoom_out               = html.Class("icon-zoom-out")
	Icon_power_off              = html.Class("icon-power-off")
	Icon_off                    = html.Class("icon-off")
	Icon_signal                 = html.Class("icon-signal")
	Icon_gear                   = html.Class("icon-gear")
	Icon_cog                    = html.Class("icon-cog")
	Icon_trash                  = html.Class("icon-trash")
	Icon_home                   = html.Class("icon-home")
	Icon_file_alt               = html.Class("icon-file-alt")
	Icon_time                   = html.Class("icon-time")
	Icon_road                   = html.Class("icon-road")
	Icon_download_alt           = html.Class("icon-download-alt")
	Icon_download               = html.Class("icon-download")
	Icon_upload                 = html.Class("icon-upload")
	Icon_inbox                  = html.Class("icon-inbox")
	Icon_play_circle            = html.Class("icon-play-circle")
	Icon_rotate_right           = html.Class("icon-rotate-right")
	Icon_repeat                 = html.Class("icon-repeat")
	Icon_refresh                = html.Class("icon-refresh")
	Icon_list_alt               = html.Class("icon-list-alt")
	Icon_lock                   = html.Class("icon-lock")
	Icon_flag                   = html.Class("icon-flag")
	Icon_headphones             = html.Class("icon-headphones")
	Icon_volume_off             = html.Class("icon-volume-off")
	Icon_volume_down            = html.Class("icon-volume-down")
	Icon_volume_up              = html.Class("icon-volume-up")
	Icon_qrcode                 = html.Class("icon-qrcode")
	Icon_barcode                = html.Class("icon-barcode")
	Icon_tag                    = html.Class("icon-tag")
	Icon_tags                   = html.Class("icon-tags")
	Icon_book                   = html.Class("icon-book")
	Icon_bookmark               = html.Class("icon-bookmark")
	Icon_print                  = html.Class("icon-print")
	Icon_camera                 = html.Class("icon-camera")
	Icon_font                   = html.Class("icon-font")
	Icon_bold                   = html.Class("icon-bold")
	Icon_italic                 = html.Class("icon-italic")
	Icon_text_height            = html.Class("icon-text-height")
	Icon_text_width             = html.Class("icon-text-width")
	Icon_align_left             = html.Class("icon-align-left")
	Icon_align_center           = html.Class("icon-align-center")
	Icon_align_right            = html.Class("icon-align-right")
	Icon_align_justify          = html.Class("icon-align-justify")
	Icon_list                   = html.Class("icon-list")
	Icon_indent_left            = html.Class("icon-indent-left")
	Icon_indent_right           = html.Class("icon-indent-right")
	Icon_facetime_video         = html.Class("icon-facetime-video")
	Icon_picture                = html.Class("icon-picture")
	Icon_pencil                 = html.Class("icon-pencil")
	Icon_map_marker             = html.Class("icon-map-marker")
	Icon_adjust                 = html.Class("icon-adjust")
	Icon_tint                   = html.Class("icon-tint")
	Icon_edit                   = html.Class("icon-edit")
	Icon_share                  = html.Class("icon-share")
	Icon_check                  = html.Class("icon-check")
	Icon_move                   = html.Class("icon-move")
	Icon_step_backward          = html.Class("icon-step-backward")
	Icon_fast_backward          = html.Class("icon-fast-backward")
	Icon_backward               = html.Class("icon-backward")
	Icon_play                   = html.Class("icon-play")
	Icon_pause                  = html.Class("icon-pause")
	Icon_stop                   = html.Class("icon-stop")
	Icon_forward                = html.Class("icon-forward")
	Icon_fast_forward           = html.Class("icon-fast-forward")
	Icon_step_forward           = html.Class("icon-step-forward")
	Icon_eject                  = html.Class("icon-eject")
	Icon_chevron_left           = html.Class("icon-chevron-left")
	Icon_chevron_right          = html.Class("icon-chevron-right")
	Icon_plus_sign              = html.Class("icon-plus-sign")
	Icon_minus_sign             = html.Class("icon-minus-sign")
	Icon_remove_sign            = html.Class("icon-remove-sign")
	Icon_ok_sign                = html.Class("icon-ok-sign")
	Icon_question_sign          = html.Class("icon-question-sign")
	Icon_info_sign              = html.Class("icon-info-sign")
	Icon_screenshot             = html.Class("icon-screenshot")
	Icon_remove_circle          = html.Class("icon-remove-circle")
	Icon_ok_circle              = html.Class("icon-ok-circle")
	Icon_ban_circle             = html.Class("icon-ban-circle")
	Icon_arrow_left             = html.Class("icon-arrow-left")
	Icon_arrow_right            = html.Class("icon-arrow-right")
	Icon_arrow_up               = html.Class("icon-arrow-up")
	Icon_arrow_down             = html.Class("icon-arrow-down")
	Icon_mail_forward           = html.Class("icon-mail-forward")
	Icon_share_alt              = html.Class("icon-share-alt")
	Icon_resize_full            = html.Class("icon-resize-full")
	Icon_resize_small           = html.Class("icon-resize-small")
	Icon_plus                   = html.Class("icon-plus")
	Icon_minus                  = html.Class("icon-minus")
	Icon_asterisk               = html.Class("icon-asterisk")
	Icon_exclamation_sign       = html.Class("icon-exclamation-sign")
	Icon_gift                   = html.Class("icon-gift")
	Icon_leaf                   = html.Class("icon-leaf")
	Icon_fire                   = html.Class("icon-fire")
	Icon_eye_open               = html.Class("icon-eye-open")
	Icon_eye_close              = html.Class("icon-eye-close")
	Icon_warning_sign           = html.Class("icon-warning-sign")
	Icon_plane                  = html.Class("icon-plane")
	Icon_calendar               = html.Class("icon-calendar")
	Icon_random                 = html.Class("icon-random")
	Icon_comment                = html.Class("icon-comment")
	Icon_magnet                 = html.Class("icon-magnet")
	Icon_chevron_up             = html.Class("icon-chevron-up")
	Icon_chevron_down           = html.Class("icon-chevron-down")
	Icon_retweet                = html.Class("icon-retweet")
	Icon_shopping_cart          = html.Class("icon-shopping-cart")
	Icon_folder_close           = html.Class("icon-folder-close")
	Icon_folder_open            = html.Class("icon-folder-open")
	Icon_resize_vertical        = html.Class("icon-resize-vertical")
	Icon_resize_horizontal      = html.Class("icon-resize-horizontal")
	Icon_bar_chart              = html.Class("icon-bar-chart")
	Icon_twitter_sign           = html.Class("icon-twitter-sign")
	Icon_facebook_sign          = html.Class("icon-facebook-sign")
	Icon_camera_retro           = html.Class("icon-camera-retro")
	Icon_key                    = html.Class("icon-key")
	Icon_gears                  = html.Class("icon-gears")
	Icon_cogs                   = html.Class("icon-cogs")
	Icon_comments               = html.Class("icon-comments")
	Icon_thumbs_up_alt          = html.Class("icon-thumbs-up-alt")
	Icon_thumbs_down_alt        = html.Class("icon-thumbs-down-alt")
	Icon_star_half              = html.Class("icon-star-half")
	Icon_heart_empty            = html.Class("icon-heart-empty")
	Icon_signout                = html.Class("icon-signout")
	Icon_linkedin_sign          = html.Class("icon-linkedin-sign")
	Icon_pushpin                = html.Class("icon-pushpin")
	Icon_external_link          = html.Class("icon-external-link")
	Icon_signin                 = html.Class("icon-signin")
	Icon_trophy                 = html.Class("icon-trophy")
	Icon_github_sign            = html.Class("icon-github-sign")
	Icon_upload_alt             = html.Class("icon-upload-alt")
	Icon_lemon                  = html.Class("icon-lemon")
	Icon_phone                  = html.Class("icon-phone")
	Icon_unchecked              = html.Class("icon-unchecked")
	Icon_check_empty            = html.Class("icon-check-empty")
	Icon_bookmark_empty         = html.Class("icon-bookmark-empty")
	Icon_phone_sign             = html.Class("icon-phone-sign")
	Icon_twitter                = html.Class("icon-twitter")
	Icon_facebook               = html.Class("icon-facebook")
	Icon_github                 = html.Class("icon-github")
	Icon_unlock                 = html.Class("icon-unlock")
	Icon_credit_card            = html.Class("icon-credit-card")
	Icon_rss                    = html.Class("icon-rss")
	Icon_hdd                    = html.Class("icon-hdd")
	Icon_bullhorn               = html.Class("icon-bullhorn")
	Icon_bell                   = html.Class("icon-bell")
	Icon_certificate            = html.Class("icon-certificate")
	Icon_hand_right             = html.Class("icon-hand-right")
	Icon_hand_left              = html.Class("icon-hand-left")
	Icon_hand_up                = html.Class("icon-hand-up")
	Icon_hand_down              = html.Class("icon-hand-down")
	Icon_circle_arrow_left      = html.Class("icon-circle-arrow-left")
	Icon_circle_arrow_right     = html.Class("icon-circle-arrow-right")
	Icon_circle_arrow_up        = html.Class("icon-circle-arrow-up")
	Icon_circle_arrow_down      = html.Class("icon-circle-arrow-down")
	Icon_globe                  = html.Class("icon-globe")
	Icon_wrench                 = html.Class("icon-wrench")
	Icon_tasks                  = html.Class("icon-tasks")
	Icon_filter                 = html.Class("icon-filter")
	Icon_briefcase              = html.Class("icon-briefcase")
	Icon_fullscreen             = html.Class("icon-fullscreen")
	Icon_group                  = html.Class("icon-group")
	Icon_link                   = html.Class("icon-link")
	Icon_cloud                  = html.Class("icon-cloud")
	Icon_beaker                 = html.Class("icon-beaker")
	Icon_cut                    = html.Class("icon-cut")
	Icon_copy                   = html.Class("icon-copy")
	Icon_paperclip              = html.Class("icon-paperclip")
	Icon_paper_clip             = html.Class("icon-paper-clip")
	Icon_save                   = html.Class("icon-save")
	Icon_sign_blank             = html.Class("icon-sign-blank")
	Icon_reorder                = html.Class("icon-reorder")
	Icon_list_ul                = html.Class("icon-list-ul")
	Icon_list_ol                = html.Class("icon-list-ol")
	Icon_strikethrough          = html.Class("icon-strikethrough")
	Icon_underline              = html.Class("icon-underline")
	Icon_table                  = html.Class("icon-table")
	Icon_magic                  = html.Class("icon-magic")
	Icon_truck                  = html.Class("icon-truck")
	Icon_pinterest              = html.Class("icon-pinterest")
	Icon_pinterest_sign         = html.Class("icon-pinterest-sign")
	Icon_google_plus_sign       = html.Class("icon-google-plus-sign")
	Icon_google_plus            = html.Class("icon-google-plus")
	Icon_money                  = html.Class("icon-money")
	Icon_caret_down             = html.Class("icon-caret-down")
	Icon_caret_up               = html.Class("icon-caret-up")
	Icon_caret_left             = html.Class("icon-caret-left")
	Icon_caret_right            = html.Class("icon-caret-right")
	Icon_columns                = html.Class("icon-columns")
	Icon_sort                   = html.Class("icon-sort")
	Icon_sort_down              = html.Class("icon-sort-down")
	Icon_sort_up                = html.Class("icon-sort-up")
	Icon_envelope               = html.Class("icon-envelope")
	Icon_linkedin               = html.Class("icon-linkedin")
	Icon_rotate_left            = html.Class("icon-rotate-left")
	Icon_undo                   = html.Class("icon-undo")
	Icon_legal                  = html.Class("icon-legal")
	Icon_dashboard              = html.Class("icon-dashboard")
	Icon_comment_alt            = html.Class("icon-comment-alt")
	Icon_comments_alt           = html.Class("icon-comments-alt")
	Icon_bolt                   = html.Class("icon-bolt")
	Icon_sitemap                = html.Class("icon-sitemap")
	Icon_umbrella               = html.Class("icon-umbrella")
	Icon_paste                  = html.Class("icon-paste")
	Icon_lightbulb              = html.Class("icon-lightbulb")
	Icon_exchange               = html.Class("icon-exchange")
	Icon_cloud_download         = html.Class("icon-cloud-download")
	Icon_cloud_upload           = html.Class("icon-cloud-upload")
	Icon_user_md                = html.Class("icon-user-md")
	Icon_stethoscope            = html.Class("icon-stethoscope")
	Icon_suitcase               = html.Class("icon-suitcase")
	Icon_bell_alt               = html.Class("icon-bell-alt")
	Icon_coffee                 = html.Class("icon-coffee")
	Icon_food                   = html.Class("icon-food")
	Icon_file_text_alt          = html.Class("icon-file-text-alt")
	Icon_building               = html.Class("icon-building")
	Icon_hospital               = html.Class("icon-hospital")
	Icon_ambulance              = html.Class("icon-ambulance")
	Icon_medkit                 = html.Class("icon-medkit")
	Icon_fighter_jet            = html.Class("icon-fighter-jet")
	Icon_beer                   = html.Class("icon-beer")
	Icon_h_sign                 = html.Class("icon-h-sign")
	Icon_plus_sign_alt          = html.Class("icon-plus-sign-alt")
	Icon_double_angle_left      = html.Class("icon-double-angle-left")
	Icon_double_angle_right     = html.Class("icon-double-angle-right")
	Icon_double_angle_up        = html.Class("icon-double-angle-up")
	Icon_double_angle_down      = html.Class("icon-double-angle-down")
	Icon_angle_left             = html.Class("icon-angle-left")
	Icon_angle_right            = html.Class("icon-angle-right")
	Icon_angle_up               = html.Class("icon-angle-up")
	Icon_angle_down             = html.Class("icon-angle-down")
	Icon_desktop                = html.Class("icon-desktop")
	Icon_laptop                 = html.Class("icon-laptop")
	Icon_tablet                 = html.Class("icon-tablet")
	Icon_mobile_phone           = html.Class("icon-mobile-phone")
	Icon_circle_blank           = html.Class("icon-circle-blank")
	Icon_quote_left             = html.Class("icon-quote-left")
	Icon_quote_right            = html.Class("icon-quote-right")
	Icon_spinner                = html.Class("icon-spinner")
	Icon_circle                 = html.Class("icon-circle")
	Icon_mail_reply             = html.Class("icon-mail-reply")
	Icon_reply                  = html.Class("icon-reply")
	Icon_github_alt             = html.Class("icon-github-alt")
	Icon_folder_close_alt       = html.Class("icon-folder-close-alt")
	Icon_folder_open_alt        = html.Class("icon-folder-open-alt")
	Icon_expand_alt             = html.Class("icon-expand-alt")
	Icon_collapse_alt           = html.Class("icon-collapse-alt")
	Icon_smile                  = html.Class("icon-smile")
	Icon_frown                  = html.Class("icon-frown")
	Icon_meh                    = html.Class("icon-meh")
	Icon_gamepad                = html.Class("icon-gamepad")
	Icon_keyboard               = html.Class("icon-keyboard")
	Icon_flag_alt               = html.Class("icon-flag-alt")
	Icon_flag_checkered         = html.Class("icon-flag-checkered")
	Icon_terminal               = html.Class("icon-terminal")
	Icon_code                   = html.Class("icon-code")
	Icon_reply_all              = html.Class("icon-reply-all")
	Icon_mail_reply_all         = html.Class("icon-mail-reply-all")
	Icon_star_half_full         = html.Class("icon-star-half-full")
	Icon_star_half_empty        = html.Class("icon-star-half-empty")
	Icon_location_arrow         = html.Class("icon-location-arrow")
	Icon_crop                   = html.Class("icon-crop")
	Icon_code_fork              = html.Class("icon-code-fork")
	Icon_unlink                 = html.Class("icon-unlink")
	Icon_question               = html.Class("icon-question")
	Icon_info                   = html.Class("icon-info")
	Icon_exclamation            = html.Class("icon-exclamation")
	Icon_superscript            = html.Class("icon-superscript")
	Icon_subscript              = html.Class("icon-subscript")
	Icon_eraser                 = html.Class("icon-eraser")
	Icon_puzzle_piece           = html.Class("icon-puzzle-piece")
	Icon_microphone             = html.Class("icon-microphone")
	Icon_microphone_off         = html.Class("icon-microphone-off")
	Icon_shield                 = html.Class("icon-shield")
	Icon_calendar_empty         = html.Class("icon-calendar-empty")
	Icon_fire_extinguisher      = html.Class("icon-fire-extinguisher")
	Icon_rocket                 = html.Class("icon-rocket")
	Icon_maxcdn                 = html.Class("icon-maxcdn")
	Icon_chevron_sign_left      = html.Class("icon-chevron-sign-left")
	Icon_chevron_sign_right     = html.Class("icon-chevron-sign-right")
	Icon_chevron_sign_up        = html.Class("icon-chevron-sign-up")
	Icon_chevron_sign_down      = html.Class("icon-chevron-sign-down")
	Icon_html5                  = html.Class("icon-html5")
	Icon_css3                   = html.Class("icon-css3")
	Icon_anchor                 = html.Class("icon-anchor")
	Icon_unlock_alt             = html.Class("icon-unlock-alt")
	Icon_bullseye               = html.Class("icon-bullseye")
	Icon_ellipsis_horizontal    = html.Class("icon-ellipsis-horizontal")
	Icon_ellipsis_vertical      = html.Class("icon-ellipsis-vertical")
	Icon_rss_sign               = html.Class("icon-rss-sign")
	Icon_play_sign              = html.Class("icon-play-sign")
	Icon_ticket                 = html.Class("icon-ticket")
	Icon_minus_sign_alt         = html.Class("icon-minus-sign-alt")
	Icon_check_minus            = html.Class("icon-check-minus")
	Icon_level_up               = html.Class("icon-level-up")
	Icon_level_down             = html.Class("icon-level-down")
	Icon_check_sign             = html.Class("icon-check-sign")
	Icon_edit_sign              = html.Class("icon-edit-sign")
	Icon_external_link_sign     = html.Class("icon-external-link-sign")
	Icon_share_sign             = html.Class("icon-share-sign")
	Icon_compass                = html.Class("icon-compass")
	Icon_collapse               = html.Class("icon-collapse")
	Icon_collapse_top           = html.Class("icon-collapse-top")
	Icon_expand                 = html.Class("icon-expand")
	Icon_euro                   = html.Class("icon-euro")
	Icon_eur                    = html.Class("icon-eur")
	Icon_gbp                    = html.Class("icon-gbp")
	Icon_dollar                 = html.Class("icon-dollar")
	Icon_usd                    = html.Class("icon-usd")
	Icon_rupee                  = html.Class("icon-rupee")
	Icon_inr                    = html.Class("icon-inr")
	Icon_yen                    = html.Class("icon-yen")
	Icon_jpy                    = html.Class("icon-jpy")
	Icon_renminbi               = html.Class("icon-renminbi")
	Icon_cny                    = html.Class("icon-cny")
	Icon_won                    = html.Class("icon-won")
	Icon_krw                    = html.Class("icon-krw")
	Icon_bitcoin                = html.Class("icon-bitcoin")
	Icon_btc                    = html.Class("icon-btc")
	Icon_file                   = html.Class("icon-file")
	Icon_file_text              = html.Class("icon-file-text")
	Icon_sort_by_alphabet       = html.Class("icon-sort-by-alphabet")
	Icon_sort_by_alphabet_alt   = html.Class("icon-sort-by-alphabet-alt")
	Icon_sort_by_attributes     = html.Class("icon-sort-by-attributes")
	Icon_sort_by_attributes_alt = html.Class("icon-sort-by-attributes-alt")
	Icon_sort_by_order          = html.Class("icon-sort-by-order")
	Icon_sort_by_order_alt      = html.Class("icon-sort-by-order-alt")
	Icon_thumbs_up              = html.Class("icon-thumbs-up")
	Icon_thumbs_down            = html.Class("icon-thumbs-down")
	Icon_youtube_sign           = html.Class("icon-youtube-sign")
	Icon_youtube                = html.Class("icon-youtube")
	Icon_xing                   = html.Class("icon-xing")
	Icon_xing_sign              = html.Class("icon-xing-sign")
	Icon_youtube_play           = html.Class("icon-youtube-play")
	Icon_dropbox                = html.Class("icon-dropbox")
	Icon_stackexchange          = html.Class("icon-stackexchange")
	Icon_instagram              = html.Class("icon-instagram")
	Icon_flickr                 = html.Class("icon-flickr")
	Icon_adn                    = html.Class("icon-adn")
	Icon_bitbucket              = html.Class("icon-bitbucket")
	Icon_bitbucket_sign         = html.Class("icon-bitbucket-sign")
	Icon_tumblr                 = html.Class("icon-tumblr")
	Icon_tumblr_sign            = html.Class("icon-tumblr-sign")
	Icon_long_arrow_down        = html.Class("icon-long-arrow-down")
	Icon_long_arrow_up          = html.Class("icon-long-arrow-up")
	Icon_long_arrow_left        = html.Class("icon-long-arrow-left")
	Icon_long_arrow_right       = html.Class("icon-long-arrow-right")
	Icon_apple                  = html.Class("icon-apple")
	Icon_windows                = html.Class("icon-windows")
	Icon_android                = html.Class("icon-android")
	Icon_linux                  = html.Class("icon-linux")
	Icon_dribbble               = html.Class("icon-dribbble")
	Icon_skype                  = html.Class("icon-skype")
	Icon_foursquare             = html.Class("icon-foursquare")
	Icon_trello                 = html.Class("icon-trello")
	Icon_female                 = html.Class("icon-female")
	Icon_male                   = html.Class("icon-male")
	Icon_gittip                 = html.Class("icon-gittip")
	Icon_sun                    = html.Class("icon-sun")
	Icon_moon                   = html.Class("icon-moon")
	Icon_archive                = html.Class("icon-archive")
	Icon_bug                    = html.Class("icon-bug")
	Icon_vk                     = html.Class("icon-vk")
	Icon_weibo                  = html.Class("icon-weibo")
	Icon_renren                 = html.Class("icon-renren")
	Ink_navigation              = html.Class("ink-navigation")
	Menu                        = html.Class("menu")
	Submenu                     = html.Class("submenu")
	Horizontal                  = html.Class("horizontal")
	Control                     = html.Class("control")
	Vertical                    = html.Class("vertical")
	Dropdown                    = html.Class("dropdown")
	Breadcrumbs                 = html.Class("breadcrumbs")
	Active                      = html.Class("active")
	Pagination                  = html.Class("pagination")
	Disabled                    = html.Class("disabled")
	Pills                       = html.Class("pills")
	Ink_dropdown                = html.Class("ink-dropdown")
	Dropdown_menu               = html.Class("dropdown-menu")
	Separator_above             = html.Class("separator-above")
	Separator_below             = html.Class("separator-below")
	Heading                     = html.Class("heading")
	White                       = html.Class("white")
	Grey                        = html.Class("grey")
	Black                       = html.Class("black")
	Orange                      = html.Class("orange")
	Blue                        = html.Class("blue")
	Green                       = html.Class("green")
	Red                         = html.Class("red")
	Flat                        = html.Class("flat")
	Rounded                     = html.Class("rounded")
	Shadowed                    = html.Class("shadowed")
	Ink_form                    = html.Class("ink-form")
	Tip                         = html.Class("tip")
	Label                       = html.Class("label")
	Input_file                  = html.Class("input-file")
	Ink_button                  = html.Class("ink-button")
	Control_group               = html.Class("control-group")
	Append_button               = html.Class("append-button")
	Append_symbol               = html.Class("append-symbol")
	Prepend_button              = html.Class("prepend-button")
	Prepend_symbol              = html.Class("prepend-symbol")
	Validation                  = html.Class("validation")
	Required                    = html.Class("required")
	Status_indicator            = html.Class("status-indicator")
	Ink_alert                   = html.Class("ink-alert")
	Basic                       = html.Class("basic")
	Block                       = html.Class("block")
	Ink_close                   = html.Class("ink-close")
	Ink_dismiss                 = html.Class("ink-dismiss")
	Ink_badge                   = html.Class("ink-badge")
	Ink_tooltip                 = html.Class("ink-tooltip")
	Content                     = html.Class("content")
	Arrow                       = html.Class("arrow")
	Up                          = html.Class("up")
	Down                        = html.Class("down")
	Left                        = html.Class("left")
	Right                       = html.Class("right")
	Ink_disabled                = html.Class("ink-disabled")
	Button_group                = html.Class("button-group")
	Button_toolbar              = html.Class("button-toolbar")
	Ink_table                   = html.Class("ink-table")
	Alternating                 = html.Class("alternating")
	Hover                       = html.Class("hover")
	Bordered                    = html.Class("bordered")
	Ink_gallery                 = html.Class("ink-gallery")
	Thumbs                      = html.Class("thumbs")
	Slider                      = html.Class("slider")
	Article_text                = html.Class("article_text")
	Example1                    = html.Class("example1")
	Example2                    = html.Class("example2")
	Stage                       = html.Class("stage")
	Next                        = html.Class("next")
	Previous                    = html.Class("previous")
	RightNav                    = html.Class("rightNav")
	Sapo_component_datepicker   = html.Class("sapo_component_datepicker")
	Sapo_cal_top_options        = html.Class("sapo_cal_top_options")
	Clean                       = html.Class("clean")
	Close                       = html.Class("close")
	Sapo_cal_top                = html.Class("sapo_cal_top")
	Sapo_cal_prev               = html.Class("sapo_cal_prev")
	Sapo_cal_next               = html.Class("sapo_cal_next")
	Sapo_cal_month_desc         = html.Class("sapo_cal_month_desc")
	Sapo_cal_month              = html.Class("sapo_cal_month")
	Sapo_cal_year_selector      = html.Class("sapo_cal_year_selector")
	Sapo_cal_month_selector     = html.Class("sapo_cal_month_selector")
	Sapo_cal_on                 = html.Class("sapo_cal_on")
	Sapo_cal_off                = html.Class("sapo_cal_off")
	Sapo_cal_header             = html.Class("sapo_cal_header")
	Sapo_cal_middle             = html.Class("sapo_cal_middle")
	Ink_modal                   = html.Class("ink-modal")
	Modal_body                  = html.Class("modal-body")
	Modal_header                = html.Class("modal-header")
	Modal_close                 = html.Class("modal-close")
	Modal_footer                = html.Class("modal-footer")
	Ink_modal_open              = html.Class("ink-modal-open")
	Ink_progress_bar            = html.Class("ink-progress-bar")
	Caption                     = html.Class("caption")
	Bar                         = html.Class("bar")
	Ink_tabs                    = html.Class("ink-tabs")
	Tabs_nav                    = html.Class("tabs-nav")
	Tabs_content                = html.Class("tabs-content")
	Top                         = html.Class("top")
	Bottom                      = html.Class("bottom")
	Ink_sortable_list           = html.Class("ink-sortable-list")
	Ink_tree_view               = html.Class("ink-tree-view")
	Open                        = html.Class("open")
	Closed                      = html.Class("closed")
	Ink_carousel                = html.Class("ink-carousel")
	Slide                       = html.Class("slide")
	Hider                       = html.Class("hider")
	Caption_over_top            = html.Class("caption-over-top")
	Caption_over_bottom         = html.Class("caption-over-bottom")
	Light                       = html.Class("light")
	Dark                        = html.Class("dark")
	Push_left                   = html.Class("push-left")
	Push_right                  = html.Class("push-right")
	Push_center                 = html.Class("push-center")
	Clearfix                    = html.Class("clearfix")
	No_margin                   = html.Class("no-margin")
	Screen_size_helper          = html.Class("screen-size-helper")
	Title                       = html.Class("title")
	Drag                        = html.Class("drag")
	Ink_shade                   = html.Class("ink-shade")
	Fade                        = html.Class("fade")
	Visible                     = html.Class("visible")
)