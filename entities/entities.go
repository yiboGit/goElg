package entities
import (
	"time"
)
type AuditUserRegister struct {
Id       int32       `db:"id" json:"id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
ConfigId       NullInt64       `db:"config_id" json:"config_id"`   
Status       int32       `db:"status" json:"status"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
TotalScore       NullInt64       `db:"total_score" json:"total_score"`   
Position       int32       `db:"position" json:"position"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
}

type BargainHelpers struct {
Id       int32       `db:"id" json:"id"`   
ActivityId       int32       `db:"activity_id" json:"activity_id"`   
FromUser       int32       `db:"from_user" json:"from_user"`   
ToUser       int32       `db:"to_user" json:"to_user"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Discount       NullFloat64       `db:"discount" json:"discount"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
}

type Bonus struct {
Id       int32       `db:"id" json:"id"`   
Code       NullString       `db:"code" json:"code"`   
App       NullString       `db:"app" json:"app"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
Money       NullFloat64       `db:"money" json:"money"`   
IsDisabled       NullInt64       `db:"is_disabled" json:"is_disabled"`   
IsReceive       NullInt64       `db:"is_receive" json:"is_receive"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsRandom       NullInt64       `db:"is_random" json:"is_random"`   
RecordId       int32       `db:"record_id" json:"record_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
Permit       NullInt64       `db:"permit" json:"permit"`   
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
IsActivate       NullInt64       `db:"is_activate" json:"is_activate"`   
NeedApproval       NullInt64       `db:"need_approval" json:"need_approval"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
ReceiveTime       NullTime       `db:"receive_time" json:"receive_time"`   // can be null, use with caution
Corp       NullString       `db:"corp" json:"corp"`   
ForTaobao       NullInt64       `db:"for_taobao" json:"for_taobao"`   
ScreenShot       NullString       `db:"screen_shot" json:"screen_shot"`   
Bind       NullInt64       `db:"bind" json:"bind"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Sn       NullString       `db:"sn" json:"sn"`   
Color       NullString       `db:"color" json:"color"`   
Extra       NullFloat64       `db:"extra" json:"extra"`   
HeadImg       NullString       `db:"head_img" json:"head_img"`   
Limit       NullInt64       `db:"limit" json:"limit"`   
NeedCheck       NullInt64       `db:"need_check" json:"need_check"`   
BillNo       NullString       `db:"bill_no" json:"bill_no"`   
Hash       NullString       `db:"hash" json:"hash"`   
Remark       NullString       `db:"remark" json:"remark"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Nickname       NullString       `db:"nickname" json:"nickname"`   
Auditor       NullString       `db:"auditor" json:"auditor"`   
ApplyTime       NullTime       `db:"apply_time" json:"apply_time"`   // can be null, use with caution
Score       NullInt64       `db:"score" json:"score"`   
ScoreStatus       NullInt64       `db:"score_status" json:"score_status"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
}

type BonusAuditor struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       NullString       `db:"name" json:"name"`   
Password       string       `db:"password" json:"password"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Provinces       NullString       `db:"provinces" json:"provinces"`   
Phone       string       `db:"phone" json:"phone"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type BonusBalance struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
Type       NullString       `db:"type" json:"type"`   
Money       NullFloat64       `db:"money" json:"money"`   
ServiceCharge       float64       `db:"service_charge" json:"service_charge"`   
IsAdd       NullInt64       `db:"is_add" json:"is_add"`   
Balance       NullFloat64       `db:"balance" json:"balance"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type BonusConfig struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
BonusName       string       `db:"bonus_name" json:"bonus_name"`   
MoneyEnabled       int32       `db:"money_enabled" json:"money_enabled"`   
ScoreEnabled       int32       `db:"score_enabled" json:"score_enabled"`   
BrandEnabled       int32       `db:"brand_enabled" json:"brand_enabled"`   
CategoryEnabled       int32       `db:"category_enabled" json:"category_enabled"`   
SnEnabled       int32       `db:"sn_enabled" json:"sn_enabled"`   
ColorEnabled       int32       `db:"color_enabled" json:"color_enabled"`   
ScreenshotEnabled       int32       `db:"screenshot_enabled" json:"screenshot_enabled"`   
SpecialBonusEnabled       int32       `db:"special_bonus_enabled" json:"special_bonus_enabled"`   
IncrBonusEnabled       int32       `db:"incr_bonus_enabled" json:"incr_bonus_enabled"`   
NameEnabled       int32       `db:"name_enabled" json:"name_enabled"`   
AreaEnabled       int32       `db:"area_enabled" json:"area_enabled"`   
StoreEnabled       int32       `db:"store_enabled" json:"store_enabled"`   
AgeEnabled       int32       `db:"age_enabled" json:"age_enabled"`   
SexEnabled       int32       `db:"sex_enabled" json:"sex_enabled"`   
AuditEnabled       int32       `db:"audit_enabled" json:"audit_enabled"`   
FixedCodeEnabled       int32       `db:"fixed_code_enabled" json:"fixed_code_enabled"`   
RandomCodeEnabled       int32       `db:"random_code_enabled" json:"random_code_enabled"`   
NeedApprovalEnabled       int32       `db:"need_approval_enabled" json:"need_approval_enabled"`   
SendAmount       int32       `db:"send_amount" json:"send_amount"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
SourceConfigId       NullInt64       `db:"source_config_id" json:"source_config_id"`   
QrCodeEnabled       int32       `db:"qr_code_enabled" json:"qr_code_enabled"`   
}

type BonusPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
TryRedirectEnabled       NullInt64       `db:"try_redirect_enabled" json:"try_redirect_enabled"`   
StatsEnabled       NullInt64       `db:"stats_enabled" json:"stats_enabled"`   
CategoryEnabled       NullInt64       `db:"category_enabled" json:"category_enabled"`   
BrandEnabled       NullInt64       `db:"brand_enabled" json:"brand_enabled"`   
SnEnabled       NullInt64       `db:"sn_enabled" json:"sn_enabled"`   
ColorEnabled       NullInt64       `db:"color_enabled" json:"color_enabled"`   
ScreenshotEnabled       NullInt64       `db:"screenshot_enabled" json:"screenshot_enabled"`   
StoreEnabled       NullInt64       `db:"store_enabled" json:"store_enabled"`   
SpecialBonusEnabled       NullInt64       `db:"special_bonus_enabled" json:"special_bonus_enabled"`   
FixedExchangeCodeEnabled       NullInt64       `db:"fixed_exchange_code_enabled" json:"fixed_exchange_code_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IncrBonusEnabled       NullInt64       `db:"incr_bonus_enabled" json:"incr_bonus_enabled"`   
AreaEnabled       NullInt64       `db:"area_enabled" json:"area_enabled"`   
NameEnabled       NullInt64       `db:"name_enabled" json:"name_enabled"`   
MainSubEnabled       NullInt64       `db:"main_sub_enabled" json:"main_sub_enabled"`   
ActStoreEnabled       NullInt64       `db:"act_store_enabled" json:"act_store_enabled"`   
ActAreaEnabled       NullInt64       `db:"act_area_enabled" json:"act_area_enabled"`   
ActNameEnabled       NullInt64       `db:"act_name_enabled" json:"act_name_enabled"`   
ActIncrEnabled       NullInt64       `db:"act_incr_enabled" json:"act_incr_enabled"`   
AuditorEnabled       NullInt64       `db:"auditor_enabled" json:"auditor_enabled"`   
AgeEnabled       NullInt64       `db:"age_enabled" json:"age_enabled"`   
SexEnabled       NullInt64       `db:"sex_enabled" json:"sex_enabled"`   
SendAmountEnabled       NullInt64       `db:"send_amount_enabled" json:"send_amount_enabled"`   
ActAuditEnabled       int32       `db:"act_audit_enabled" json:"act_audit_enabled"`   
AuditEnabled       int32       `db:"audit_enabled" json:"audit_enabled"`   
ActScreenshotEnabled       NullInt64       `db:"act_screenshot_enabled" json:"act_screenshot_enabled"`   
MoneyEnabled       NullInt64       `db:"money_enabled" json:"money_enabled"`   
ScoreEnabled       NullInt64       `db:"score_enabled" json:"score_enabled"`   
WechatOpenEnabled       NullInt64       `db:"wechat_open_enabled" json:"wechat_open_enabled"`   
IsUpstream       NullInt64       `db:"is_upstream" json:"is_upstream"`   
IsDownstream       NullInt64       `db:"is_downstream" json:"is_downstream"`   
TryRedirectTime       NullString       `db:"try_redirect_time" json:"try_redirect_time"`   
StatsTime       NullString       `db:"stats_time" json:"stats_time"`   
AuditorTime       NullString       `db:"auditor_time" json:"auditor_time"`   
WechatOpenTime       NullString       `db:"wechat_open_time" json:"wechat_open_time"`   
IsUpperTime       NullString       `db:"is_upper_time" json:"is_upper_time"`   
IsLowerTime       NullString       `db:"is_lower_time" json:"is_lower_time"`   
SourceId       NullInt64       `db:"source_id" json:"source_id"`   
BeginTime       NullString       `db:"begin_time" json:"begin_time"`   
}

type BonusRecord struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
TotalMoney       NullFloat64       `db:"total_money" json:"total_money"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsActivate       NullInt64       `db:"is_activate" json:"is_activate"`   
NeedApproval       NullInt64       `db:"need_approval" json:"need_approval"`   
IsRandom       NullInt64       `db:"is_random" json:"is_random"`   
ForTaobao       NullInt64       `db:"for_taobao" json:"for_taobao"`   
FixedCode       NullString       `db:"fixed_code" json:"fixed_code"`   
Bind       NullInt64       `db:"bind" json:"bind"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Sn       NullString       `db:"sn" json:"sn"`   
Color       NullString       `db:"color" json:"color"`   
Specials       NullString       `db:"specials" json:"specials"`   
SendOne       NullInt64       `db:"send_one" json:"send_one"`   
Increments       NullString       `db:"increments" json:"increments"`   
Name       NullString       `db:"name" json:"name"`   
Remark       NullString       `db:"remark" json:"remark"`   
SendAmount       int32       `db:"send_amount" json:"send_amount"`   
Score       NullInt64       `db:"score" json:"score"`   
ConfigId       int32       `db:"config_id" json:"config_id"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
}

type BonusShop struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
IsMoney       NullInt64       `db:"is_money" json:"is_money"`   
Name       NullString       `db:"name" json:"name"`   
Img       NullString       `db:"img" json:"img"`   
Money       NullFloat64       `db:"money" json:"money"`   
CostScore       int32       `db:"cost_score" json:"cost_score"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
OnUse       NullInt64       `db:"on_use" json:"on_use"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
DetailImg       NullString       `db:"detail_img" json:"detail_img"`   
DetailDesc       NullString       `db:"detail_desc" json:"detail_desc"`   
}

type BonusUser struct {
Id       int32       `db:"id" json:"id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TotalScore       NullInt64       `db:"total_score" json:"total_score"`   
UserPhone       string       `db:"user_phone" json:"user_phone"`   
Sex       NullString       `db:"sex" json:"sex"`   
Corp       NullString       `db:"corp" json:"corp"`   
HeadImg       NullString       `db:"head_img" json:"head_img"`   
Name       NullString       `db:"name" json:"name"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Nickname       NullString       `db:"nickname" json:"nickname"`   
Age       NullString       `db:"age" json:"age"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type BrandImages struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
Img4       NullString       `db:"img_4" json:"img_4"`   
}

type CategoryImages struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
Category       string       `db:"category" json:"category"`   
}

type EAccount struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Type       NullString       `db:"type" json:"type"`   
Name       NullString       `db:"name" json:"name"`   
Manager       NullInt64       `db:"manager" json:"manager"`   
Bank       NullString       `db:"bank" json:"bank"`   
CardNum       NullInt64       `db:"card_num" json:"card_num"`   
Balance       NullFloat64       `db:"balance" json:"balance"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EActivity struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       NullString       `db:"name" json:"name"`   
QrConfig       NullString       `db:"qr_config" json:"qr_config"`   
Images       NullString       `db:"images" json:"images"`   
BeginTime       NullTime       `db:"begin_time" json:"begin_time"`   // can be null, use with caution
EndTime       NullTime       `db:"end_time" json:"end_time"`   // can be null, use with caution
Rules       NullString       `db:"rules" json:"rules"`   
SubscribeMsg       NullString       `db:"subscribe_msg" json:"subscribe_msg"`   
Type       NullString       `db:"type" json:"type"`   
Keyword       NullString       `db:"keyword" json:"keyword"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Money       NullInt64       `db:"money" json:"money"`   
Minus       NullInt64       `db:"minus" json:"minus"`   
}

type EAgent struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Granted       int32       `db:"granted" json:"granted"`   
Account       string       `db:"account" json:"account"`   
Password       string       `db:"password" json:"password"`   
Phone       string       `db:"phone" json:"phone"`   
Name       string       `db:"name" json:"name"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EAnnounce struct {
Id       int32       `db:"id" json:"id"`   
Title       NullString       `db:"title" json:"title"`   
Content       NullString       `db:"content" json:"content"`   
ExpiredTime       string       `db:"expired_time" json:"expired_time"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
OnUse       NullInt64       `db:"on_use" json:"on_use"`   
ScrmEnabled       NullInt64       `db:"scrm_enabled" json:"scrm_enabled"`   
JxeEnabled       NullInt64       `db:"jxe_enabled" json:"jxe_enabled"`   
Subtitle       NullString       `db:"subtitle" json:"subtitle"`   
}

type EBasicGoods struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Category       string       `db:"category" json:"category"`   
Sn       string       `db:"sn" json:"sn"`   
Brand       string       `db:"brand" json:"brand"`   
Series       string       `db:"series" json:"series"`   
Model       NullString       `db:"model" json:"model"`   
Refract       NullString       `db:"refract" json:"refract"`   
Color       NullString       `db:"color" json:"color"`   
Sph       NullFloat64       `db:"sph" json:"sph"`   
Cyl       NullFloat64       `db:"cyl" json:"cyl"`   
IsAdd       NullInt64       `db:"is_add" json:"is_add"`   
Storage       NullInt64       `db:"storage" json:"storage"`   
RetailPrice       NullFloat64       `db:"retail_price" json:"retail_price"`   
TradePrice       NullFloat64       `db:"trade_price" json:"trade_price"`   
CostPrice       NullFloat64       `db:"cost_price" json:"cost_price"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Add       NullFloat64       `db:"add" json:"add"`   
ModelSubName       NullString       `db:"model_sub_name" json:"model_sub_name"`   
}

type EBatch struct {
Id       int32       `db:"id" json:"id"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
Title       string       `db:"title" json:"title"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EBatchGoods struct {
Id       int32       `db:"id" json:"id"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
EGoodsId       int32       `db:"e_goods_id" json:"e_goods_id"`   
EBatchId       int32       `db:"e_batch_id" json:"e_batch_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type ECardActivity struct {
Id       int32       `db:"id" json:"id"`   
Openid       string       `db:"openid" json:"openid"`   
CardId       string       `db:"card_id" json:"card_id"`   
Appid       string       `db:"appid" json:"appid"`   
ActivityId       int32       `db:"activity_id" json:"activity_id"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
}

type EChannel struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
SceneId       int32       `db:"scene_id" json:"scene_id"`   
Url       NullString       `db:"url" json:"url"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
ChannelName       NullString       `db:"channel_name" json:"channel_name"`   
ActionName       NullString       `db:"action_name" json:"action_name"`   
ExpireSeconds       NullInt64       `db:"expire_seconds" json:"expire_seconds"`   
IsWelcome       NullInt64       `db:"is_welcome" json:"is_welcome"`   
WelcomeMsg       NullString       `db:"welcome_msg" json:"welcome_msg"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
ChannelTicket       NullString       `db:"channel_ticket" json:"channel_ticket"`   
Type       int32       `db:"type" json:"type"`   
Remark       NullString       `db:"remark" json:"remark"`   
ScanNum       NullInt64       `db:"scan_num" json:"scan_num"`   
SubscribeNum       NullInt64       `db:"subscribe_num" json:"subscribe_num"`   
UnsubscribeNum       NullInt64       `db:"unsubscribe_num" json:"unsubscribe_num"`   
MediaId       NullString       `db:"media_id" json:"media_id"`   
ActivityId       NullInt64       `db:"activity_id" json:"activity_id"`   
SubscribeMsg       NullString       `db:"subscribe_msg" json:"subscribe_msg"`   
ActStaffId       NullInt64       `db:"act_staff_id" json:"act_staff_id"`   
ActStoreId       NullInt64       `db:"act_store_id" json:"act_store_id"`   
BargainUrl       NullString       `db:"bargain_url" json:"bargain_url"`   
BelongedStoreId       NullInt64       `db:"belonged_store_id" json:"belonged_store_id"`   
PicUrl       NullString       `db:"pic_url" json:"pic_url"`   
ShareRange       NullString       `db:"share_range" json:"share_range"`   
}

type EChannelUser struct {
Id       int32       `db:"id" json:"id"`   
ChannelId       int32       `db:"channel_id" json:"channel_id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Event       NullString       `db:"event" json:"event"`   
EventKey       NullString       `db:"event_key" json:"event_key"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EChatMaterial struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Type       string       `db:"type" json:"type"`   
Content       NullString       `db:"content" json:"content"`   
ImageTitle       NullString       `db:"image_title" json:"image_title"`   
ImageUrl       NullString       `db:"image_url" json:"image_url"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
NewsId       NullInt64       `db:"news_id" json:"news_id"`   
CouponId       NullInt64       `db:"coupon_id" json:"coupon_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
MediaId       NullString       `db:"media_id" json:"media_id"`   
ExpiredTime       NullInt64       `db:"expired_time" json:"expired_time"`   
}

type EChatMessage struct {
Id       int32       `db:"id" json:"id"`   
Openid       string       `db:"openid" json:"openid"`   
SendTime       int32       `db:"send_time" json:"send_time"`   
Content        string       `db:"content" json:"content"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
SessionId       int32       `db:"session_id" json:"session_id"`   
Appid       string       `db:"appid" json:"appid"`   
MsgType       NullString       `db:"msg_type" json:"msg_type"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
MaterialId       NullInt64       `db:"material_id" json:"material_id"`   
ChannelId       NullInt64       `db:"channel_id" json:"channel_id"`   
FromStaff       NullInt64       `db:"from_staff" json:"from_staff"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
Duration       NullInt64       `db:"duration" json:"duration"`   
UserId       int32       `db:"user_id" json:"user_id"`   
TemplateId       NullInt64       `db:"template_id" json:"template_id"`   
Result       NullString       `db:"result" json:"result"`   
}

type EChatSession struct {
Id       int32       `db:"id" json:"id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
Openid       string       `db:"openid" json:"openid"`   
Appid       string       `db:"appid" json:"appid"`   
LastActive       int32       `db:"last_active" json:"last_active"`   
UserLastActive       NullInt64       `db:"user_last_active" json:"user_last_active"`   
}

type EComment struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
OptometryId       NullInt64       `db:"optometry_id" json:"optometry_id"`   
ConsumptionId       NullInt64       `db:"consumption_id" json:"consumption_id"`   
Grade       NullFloat64       `db:"grade" json:"grade"`   
Remark       NullString       `db:"remark" json:"remark"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EConsult struct {
Id       int32       `db:"id" json:"id"`   
Question       NullString       `db:"question" json:"question"`   
Answer       NullString       `db:"answer" json:"answer"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
}

type EConsum struct {
Id       int32       `db:"id" json:"id"`   
MemberId       int32       `db:"member_id" json:"member_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Sn       NullString       `db:"sn" json:"sn"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
GoodsInfo       NullString       `db:"goods_info" json:"goods_info"`   
Num       NullInt64       `db:"num" json:"num"`   
OrderPrice       NullFloat64       `db:"order_price" json:"order_price"`   
Price       float64       `db:"price" json:"price"`   
Source       string       `db:"source" json:"source"`   
Status       string       `db:"status" json:"status"`   
Remark       string       `db:"remark" json:"remark"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
PayMethod       NullString       `db:"pay_method" json:"pay_method"`   
}

type EConsumItems struct {
Id       int32       `db:"id" json:"id"`   
ConsumptionId       int32       `db:"consumption_id" json:"consumption_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
GoodsName       NullString       `db:"goods_name" json:"goods_name"`   
Num       NullInt64       `db:"num" json:"num"`   
Price       NullFloat64       `db:"price" json:"price"`   
Unit       NullString       `db:"unit" json:"unit"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EConsumption struct {
Id       int32       `db:"id" json:"id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
OrderSn       NullString       `db:"order_sn" json:"order_sn"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
GoodsInfo       NullString       `db:"goods_info" json:"goods_info"`   
Num       NullInt64       `db:"num" json:"num"`   
Price       float64       `db:"price" json:"price"`   
DiscountPrice       float64       `db:"discount_price" json:"discount_price"`   
DiscountScore       int32       `db:"discount_score" json:"discount_score"`   
BackScore       int32       `db:"back_score" json:"back_score"`   
Source       string       `db:"source" json:"source"`   
OrderStatus       string       `db:"order_status" json:"order_status"`   
CostumerRemark       NullString       `db:"costumer_remark" json:"costumer_remark"`   
Remark       string       `db:"remark" json:"remark"`   
TryFavorScoreRatio       NullInt64       `db:"try_favor_score_ratio" json:"try_favor_score_ratio"`   
TryReduceMoney       NullFloat64       `db:"try_reduce_money" json:"try_reduce_money"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Disabled       NullInt64       `db:"disabled" json:"disabled"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
CardPrice       NullFloat64       `db:"card_price" json:"card_price"`   
CustomerRemark       NullString       `db:"customer_remark" json:"customer_remark"`   
PayMethod       NullString       `db:"pay_method" json:"pay_method"`   
Relation       NullString       `db:"relation" json:"relation"`   
YxdId       NullString       `db:"yxd_id" json:"yxd_id"`   
ExcludePrice       float64       `db:"exclude_price" json:"exclude_price"`   
CreatorId       NullInt64       `db:"creator_id" json:"creator_id"`   
}

type EConsumptionGoods struct {
Id       int32       `db:"id" json:"id"`   
ConsumptionId       int32       `db:"consumption_id" json:"consumption_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
GoodsName       NullString       `db:"goods_name" json:"goods_name"`   
GoodsInfo       NullString       `db:"goods_info" json:"goods_info"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Sn       NullString       `db:"sn" json:"sn"`   
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
ColorSn       NullString       `db:"color_sn" json:"color_sn"`   
Num       NullInt64       `db:"num" json:"num"`   
Price       NullFloat64       `db:"price" json:"price"`   
Unit       NullString       `db:"unit" json:"unit"`   
SocialLikeNum       NullInt64       `db:"social_like_num" json:"social_like_num"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type ECoupon struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CardType       NullString       `db:"card_type" json:"card_type"`   
DateType       NullString       `db:"date_type" json:"date_type"`   
CardId       NullString       `db:"card_id" json:"card_id"`   
Title       NullString       `db:"title" json:"title"`   
InitialStocks       NullInt64       `db:"initial_stocks" json:"initial_stocks"`   
CurrentStocks       NullInt64       `db:"current_stocks" json:"current_stocks"`   
Status       NullString       `db:"status" json:"status"`   
ConsumeWay       NullString       `db:"consume_way" json:"consume_way"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
BeginTimestamp       NullTime       `db:"begin_timestamp" json:"begin_timestamp"`   // can be null, use with caution
EndTimestamp       NullTime       `db:"end_timestamp" json:"end_timestamp"`   // can be null, use with caution
FixedTerm       NullInt64       `db:"fixed_term" json:"fixed_term"`   
FixedBeginTerm       NullInt64       `db:"fixed_begin_term" json:"fixed_begin_term"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CardImg       NullString       `db:"card_img" json:"card_img"`   
}

type ECouponCosume struct {
Id       int32       `db:"id" json:"id"`   
CouponId       int32       `db:"coupon_id" json:"coupon_id"`   
ActivityId       NullInt64       `db:"activity_id" json:"activity_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
WxId       string       `db:"wx_id" json:"wx_id"`   
CouponTitle       NullString       `db:"coupon_title" json:"coupon_title"`   
Code       NullString       `db:"code" json:"code"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
StoreName       NullString       `db:"store_name" json:"store_name"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
StaffName       NullString       `db:"staff_name" json:"staff_name"`   
LocationName       NullString       `db:"location_name" json:"location_name"`   
StaffOpenid       NullString       `db:"staff_openid" json:"staff_openid"`   
VerifyCode       NullString       `db:"verify_code" json:"verify_code"`   
RemarkAmount       NullString       `db:"remark_amount" json:"remark_amount"`   
CheckConsume       NullInt64       `db:"check_consume" json:"check_consume"`   
ConsumeSource       NullString       `db:"consume_source" json:"consume_source"`   
ConsumeWay       NullString       `db:"consume_way" json:"consume_way"`   
ConsumeTime       NullTime       `db:"consume_time" json:"consume_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ActivityName       NullString       `db:"activity_name" json:"activity_name"`   
CardId       NullString       `db:"card_id" json:"card_id"`   
}

type ECouponStore struct {
Id       int32       `db:"id" json:"id"`   
CouponId       int32       `db:"coupon_id" json:"coupon_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type ECourse struct {
Id       int32       `db:"id" json:"id"`   
Title       NullString       `db:"title" json:"title"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
Text       NullString       `db:"text" json:"text"`   
SubTitle       NullString       `db:"sub_title" json:"sub_title"`   
Thumbing       NullString       `db:"thumbing" json:"thumbing"`   
Category       NullString       `db:"category" json:"category"`   
Type       NullString       `db:"type" json:"type"`   
Cover       NullString       `db:"cover" json:"cover"`   
OnUse       NullInt64       `db:"on_use" json:"on_use"`   
}

type EFinacialDoc struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Sn       string       `db:"sn" json:"sn"`   
OperatorId       NullInt64       `db:"operator_id" json:"operator_id"`   
OperatorName       NullString       `db:"operator_name" json:"operator_name"`   
SendId       NullInt64       `db:"send_id" json:"send_id"`   
SendName       NullString       `db:"send_name" json:"send_name"`   
ReceiverId       NullInt64       `db:"receiver_id" json:"receiver_id"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
Type       string       `db:"type" json:"type"`   
SubType       string       `db:"sub_type" json:"sub_type"`   
OptometryUserId       NullInt64       `db:"optometry_user_id" json:"optometry_user_id"`   
ConsumptionId       NullInt64       `db:"consumption_id" json:"consumption_id"`   
AccountId       NullInt64       `db:"account_id" json:"account_id"`   
ReceivableMoney       NullFloat64       `db:"receivable_money" json:"receivable_money"`   
ActualMoney       NullFloat64       `db:"actual_money" json:"actual_money"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EGoods struct {
Id       int32       `db:"id" json:"id"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
Brand       string       `db:"brand" json:"brand"`   
Cat       string       `db:"cat" json:"cat"`   
Model       string       `db:"model" json:"model"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EGoodsCategory struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Img       NullString       `db:"img" json:"img"`   
Description       NullString       `db:"description" json:"description"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EGoodsImage struct {
Id       int32       `db:"id" json:"id"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
EGoodsId       int32       `db:"e_goods_id" json:"e_goods_id"`   
EBatchId       int32       `db:"e_batch_id" json:"e_batch_id"`   
ImgUrl       string       `db:"img_url" json:"img_url"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EGoodsLib struct {
Id       int32       `db:"id" json:"id"`   
CategoryId       int32       `db:"category_id" json:"category_id"`   
Title       string       `db:"title" json:"title"`   
Stock       int32       `db:"stock" json:"stock"`   
Price       float64       `db:"price" json:"price"`   
IsShow       int32       `db:"is_show" json:"is_show"`   
Thumbnail       string       `db:"thumbnail" json:"thumbnail"`   
Banner       NullString       `db:"banner" json:"banner"`   
Detail       NullString       `db:"detail" json:"detail"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type ELocation struct {
Id       int32       `db:"id" json:"id"`   
WarehouseId       int32       `db:"warehouse_id" json:"warehouse_id"`   
Name       string       `db:"name" json:"name"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EMember struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Name       string       `db:"name" json:"name"`   
Sex       NullString       `db:"sex" json:"sex"`   
Birthday       NullString       `db:"birthday" json:"birthday"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ThirdId       int32       `db:"third_id" json:"third_id"`   
AddressId       NullInt64       `db:"address_id" json:"address_id"`   
OptometryId       NullInt64       `db:"optometry_id" json:"optometry_id"`   
HeadUrl       NullString       `db:"head_url" json:"head_url"`   
YxdUsers       NullString       `db:"yxd_users" json:"yxd_users"`   
Sync       NullInt64       `db:"sync" json:"sync"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
ChannelId       NullInt64       `db:"channel_id" json:"channel_id"`   
IsSubscribe       NullInt64       `db:"is_subscribe" json:"is_subscribe"`   
SubscribeTime       NullTime       `db:"subscribe_time" json:"subscribe_time"`   // can be null, use with caution
Score       NullInt64       `db:"score" json:"score"`   
Country       NullString       `db:"country" json:"country"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
}

type ENotfiyUser struct {
Id       int32       `db:"id" json:"id"`   
WxId       string       `db:"wx_id" json:"wx_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
NotifyRecordId       int32       `db:"notify_record_id" json:"notify_record_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
UserPhone       NullString       `db:"user_phone" json:"user_phone"`   
Status       string       `db:"status" json:"status"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type ENotifyMessage struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
TemplateId       NullInt64       `db:"template_id" json:"template_id"`   
Type       NullString       `db:"type" json:"type"`   
ContentList       NullString       `db:"content_list" json:"content_list"`   
Result       NullString       `db:"result" json:"result"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Content       NullString       `db:"content" json:"content"`   
}

type ENotifyRecord struct {
Id       int32       `db:"id" json:"id"`   
Owner       int32       `db:"owner" json:"owner"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Sex       NullString       `db:"sex" json:"sex"`   
ChannelIds       NullString       `db:"channel_ids" json:"channel_ids"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
MinScore       NullInt64       `db:"min_score" json:"min_score"`   
MaxScore       NullInt64       `db:"max_score" json:"max_score"`   
NotifyTemplateId       NullInt64       `db:"notify_template_id" json:"notify_template_id"`   
NotifyContent       NullString       `db:"notify_content" json:"notify_content"`   
NotifyUrl       NullString       `db:"notify_url" json:"notify_url"`   
TimeInterval       NullString       `db:"time_interval" json:"time_interval"`   
TransferNum       NullInt64       `db:"transfer_num" json:"transfer_num"`   
TagIds       NullString       `db:"tag_ids" json:"tag_ids"`   
Name       NullString       `db:"name" json:"name"`   
}

type ENotifyTemplate struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
MsgId       NullString       `db:"msg_id" json:"msg_id"`   
Name       NullString       `db:"name" json:"name"`   
Title       NullString       `db:"title" json:"title"`   
ContentList       NullString       `db:"content_list" json:"content_list"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Link       NullString       `db:"link" json:"link"`   
CycleTime       NullInt64       `db:"cycle_time" json:"cycle_time"`   
Type       NullString       `db:"type" json:"type"`   
}

type EOptometry struct {
Id       int32       `db:"id" json:"id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
SphLeft       float64       `db:"sph_left" json:"sph_left"`   
SphRight       float64       `db:"sph_right" json:"sph_right"`   
CylLeft       NullFloat64       `db:"cyl_left" json:"cyl_left"`   
CylRight       NullFloat64       `db:"cyl_right" json:"cyl_right"`   
AxisLeft       NullInt64       `db:"axis_left" json:"axis_left"`   
AxisRight       NullInt64       `db:"axis_right" json:"axis_right"`   
Pd       float64       `db:"pd" json:"pd"`   
Remark       NullString       `db:"remark" json:"remark"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
CorctVisionLeft       NullFloat64       `db:"corct_vision_left" json:"corct_vision_left"`   
CorctVisionRight       NullFloat64       `db:"corct_vision_right" json:"corct_vision_right"`   
PdLeft       NullFloat64       `db:"pd_left" json:"pd_left"`   
PdRight       NullFloat64       `db:"pd_right" json:"pd_right"`   
Ph       NullInt64       `db:"ph" json:"ph"`   
OriginVisionLeft       NullFloat64       `db:"origin_vision_left" json:"origin_vision_left"`   
OriginVisionRight       NullFloat64       `db:"origin_vision_right" json:"origin_vision_right"`   
PrismLeft       NullInt64       `db:"prism_left" json:"prism_left"`   
PrismRight       NullInt64       `db:"prism_right" json:"prism_right"`   
StatusLeft       NullString       `db:"status_left" json:"status_left"`   
StatusRight       NullString       `db:"status_right" json:"status_right"`   
AddLeft       NullFloat64       `db:"add_left" json:"add_left"`   
AddRight       NullFloat64       `db:"add_right" json:"add_right"`   
Source       NullString       `db:"source" json:"source"`   
Relation       string       `db:"relation" json:"relation"`   
Sn       NullString       `db:"sn" json:"sn"`   
UsedFor       NullString       `db:"used_for" json:"used_for"`   
DomainEye       NullString       `db:"domain_eye" json:"domain_eye"`   
PhLeft       NullInt64       `db:"ph_left" json:"ph_left"`   
PhRight       NullInt64       `db:"ph_right" json:"ph_right"`   
YxdId       NullString       `db:"yxd_id" json:"yxd_id"`   
IsDetail       NullInt64       `db:"is_detail" json:"is_detail"`   
IsBasic       NullInt64       `db:"is_basic" json:"is_basic"`   
CreatorId       NullInt64       `db:"creator_id" json:"creator_id"`   
PrismIoLeft       NullInt64       `db:"prism_io_left" json:"prism_io_left"`   
PrismIoRight       NullInt64       `db:"prism_io_right" json:"prism_io_right"`   
PrismUdLeft       NullInt64       `db:"prism_ud_left" json:"prism_ud_left"`   
PrismUdRight       NullInt64       `db:"prism_ud_right" json:"prism_ud_right"`   
PrismIoDirectionLeft       NullString       `db:"prism_io_direction_left" json:"prism_io_direction_left"`   
PrismIoDirectionRight       NullString       `db:"prism_io_direction_right" json:"prism_io_direction_right"`   
PrismUdDirectionLeft       NullString       `db:"prism_ud_direction_left" json:"prism_ud_direction_left"`   
PrismUdDirectionRight       NullString       `db:"prism_ud_direction_right" json:"prism_ud_direction_right"`   
OptometryImg       NullString       `db:"optometry_img" json:"optometry_img"`   
}

type EOptometryUser struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
PhoneWx       NullString       `db:"phone_wx" json:"phone_wx"`   
Name       string       `db:"name" json:"name"`   
Sex       NullString       `db:"sex" json:"sex"`   
Birthday       NullString       `db:"birthday" json:"birthday"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ThirdId       int32       `db:"third_id" json:"third_id"`   
AddressId       NullInt64       `db:"address_id" json:"address_id"`   
OptometryId       NullInt64       `db:"optometry_id" json:"optometry_id"`   
HeadUrl       NullString       `db:"head_url" json:"head_url"`   
YxdUsers       NullString       `db:"yxd_users" json:"yxd_users"`   
Sync       NullInt64       `db:"sync" json:"sync"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
ChannelId       NullInt64       `db:"channel_id" json:"channel_id"`   
IsSubscribe       NullInt64       `db:"is_subscribe" json:"is_subscribe"`   
SubscribeTime       NullTime       `db:"subscribe_time" json:"subscribe_time"`   // can be null, use with caution
Score       NullInt64       `db:"score" json:"score"`   
Country       NullString       `db:"country" json:"country"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
SocialUserId       NullInt64       `db:"social_user_id" json:"social_user_id"`   
Unionid       NullString       `db:"unionid" json:"unionid"`   
UnsubscribeTime       NullTime       `db:"unsubscribe_time" json:"unsubscribe_time"`   // can be null, use with caution
}

type EPay struct {
Id       int32       `db:"id" json:"id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
GoodsType       string       `db:"goods_type" json:"goods_type"`   
GoodsDetail       string       `db:"goods_detail" json:"goods_detail"`   
Money       float64       `db:"money" json:"money"`   
Status       int32       `db:"status" json:"status"`   
TradeNo       string       `db:"trade_no" json:"trade_no"`   
Hash       string       `db:"hash" json:"hash"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EPosition struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
PositionName       string       `db:"position_name" json:"position_name"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
AgentId       NullInt64       `db:"agent_id" json:"agent_id"`   
}

type EPrinter struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
Sn       string       `db:"sn" json:"sn"`   
Key       string       `db:"key" json:"key"`   
Name       NullString       `db:"name" json:"name"`   
TitleList       NullString       `db:"title_list" json:"title_list"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CenterId       NullInt64       `db:"center_id" json:"center_id"`   
}

type EPurchaseDoc struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Sn       string       `db:"sn" json:"sn"`   
State       NullString       `db:"state" json:"state"`   
OperatorId       NullInt64       `db:"operator_id" json:"operator_id"`   
OperatorName       NullString       `db:"operator_name" json:"operator_name"`   
ApplierId       NullInt64       `db:"applier_id" json:"applier_id"`   
ApplierName       NullString       `db:"applier_name" json:"applier_name"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Total       NullInt64       `db:"total" json:"total"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EPurchaseDocItems struct {
Id       int32       `db:"id" json:"id"`   
Goods       NullString       `db:"goods" json:"goods"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Num       NullInt64       `db:"num" json:"num"`   
DocId       NullInt64       `db:"doc_id" json:"doc_id"`   
Category       NullString       `db:"category" json:"category"`   
}

type ERemindRecord struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
OptometryId       int32       `db:"optometry_id" json:"optometry_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
Type       NullString       `db:"type" json:"type"`   
RemindTime       NullTime       `db:"remind_time" json:"remind_time"`   // can be null, use with caution
Status       NullString       `db:"status" json:"status"`   
NotifyTemplateId       NullInt64       `db:"notify_template_id" json:"notify_template_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Feedback       NullInt64       `db:"feedback" json:"feedback"`   
}

type ESaleGoods struct {
Id       int32       `db:"id" json:"id"`   
ConsumptionId       int32       `db:"consumption_id" json:"consumption_id"`   
TypeId       int32       `db:"type_id" json:"type_id"`   
Type       NullString       `db:"type" json:"type"`   
Name       NullString       `db:"name" json:"name"`   
Number       NullInt64       `db:"number" json:"number"`   
Price       NullInt64       `db:"price" json:"price"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type ESaleType struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TypeName       NullString       `db:"type_name" json:"type_name"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EScore struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
ConsumptionId       NullInt64       `db:"consumption_id" json:"consumption_id"`   
Source       string       `db:"source" json:"source"`   
Type       string       `db:"type" json:"type"`   
Score       int32       `db:"score" json:"score"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Disabled       NullInt64       `db:"disabled" json:"disabled"`   
ScoreOrderId       NullInt64       `db:"score_order_id" json:"score_order_id"`   
}

type EScoreGoods struct {
Id       int32       `db:"id" json:"id"`   
Type       int32       `db:"type" json:"type"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CategoryId       int32       `db:"category_id" json:"category_id"`   
Title       string       `db:"title" json:"title"`   
Thumbnail       NullString       `db:"thumbnail" json:"thumbnail"`   
Banner       NullString       `db:"banner" json:"banner"`   
Img       NullString       `db:"img" json:"img"`   
Detail       NullString       `db:"detail" json:"detail"`   
GoodsLibId       NullInt64       `db:"goods_lib_id" json:"goods_lib_id"`   
Description       NullString       `db:"description" json:"description"`   
Rank       int32       `db:"rank" json:"rank"`   
CouponId       NullInt64       `db:"coupon_id" json:"coupon_id"`   
Stock       int32       `db:"stock" json:"stock"`   
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
SellingPrice       NullFloat64       `db:"selling_price" json:"selling_price"`   
Score       NullInt64       `db:"score" json:"score"`   
IsShow       int32       `db:"is_show" json:"is_show"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
IsRecommend       int32       `db:"is_recommend" json:"is_recommend"`   
IsOriginal       int32       `db:"is_original" json:"is_original"`   
}

type EScoreLayout struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Banner       NullString       `db:"banner" json:"banner"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EScoreOrder struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
Sn       string       `db:"sn" json:"sn"`   
ScoreAmount       NullInt64       `db:"score_amount" json:"score_amount"`   
MoneyAmount       NullFloat64       `db:"money_amount" json:"money_amount"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
Address       NullString       `db:"address" json:"address"`   
AddressId       NullInt64       `db:"address_id" json:"address_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
PayTime       NullTime       `db:"pay_time" json:"pay_time"`   // can be null, use with caution
CheckStatus       NullString       `db:"check_status" json:"check_status"`   
Source       NullString       `db:"source" json:"source"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
TradeNo       NullString       `db:"tradeNo" json:"tradeNo"`   
}

type EScoreOrderGoods struct {
Id       int32       `db:"id" json:"id"`   
ScoreOrderId       int32       `db:"score_order_id" json:"score_order_id"`   
CategoryId       int32       `db:"category_id" json:"category_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Title       NullString       `db:"title" json:"title"`   
Type       NullInt64       `db:"type" json:"type"`   
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
Score       NullInt64       `db:"score" json:"score"`   
SellingPrice       NullFloat64       `db:"selling_price" json:"selling_price"`   
Thumbnail       NullString       `db:"thumbnail" json:"thumbnail"`   
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
CheckStaffId       NullInt64       `db:"check_staff_id" json:"check_staff_id"`   
CheckStatus       NullString       `db:"check_status" json:"check_status"`   
CheckStoreId       NullInt64       `db:"check_store_id" json:"check_store_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CheckTime       NullTime       `db:"check_time" json:"check_time"`   // can be null, use with caution
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
PayTime       NullTime       `db:"pay_time" json:"pay_time"`   // can be null, use with caution
UseOriginal       int32       `db:"use_original" json:"use_original"`   
}

type EShopCart struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
Selected       NullInt64       `db:"selected" json:"selected"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type EShoppingQueue struct {
Id       int32       `db:"id" json:"id"`   
Day       NullTime       `db:"day" json:"day"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
UserId       int32       `db:"user_id" json:"user_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
}

type ESource struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Num       int32       `db:"num" json:"num"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EStaff struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Granted       int32       `db:"granted" json:"granted"`   
Account       string       `db:"account" json:"account"`   
Password       string       `db:"password" json:"password"`   
Phone       string       `db:"phone" json:"phone"`   
Name       string       `db:"name" json:"name"`   
Position       string       `db:"position" json:"position"`   
BelongsTo       string       `db:"belongs_to" json:"belongs_to"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Deleted       int32       `db:"deleted" json:"deleted"`   
IsThird       NullInt64       `db:"is_third" json:"is_third"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
OrderRecordEnabled       NullInt64       `db:"order_record_enabled" json:"order_record_enabled"`   
OptometryDataEnabled       NullInt64       `db:"optometry_data_enabled" json:"optometry_data_enabled"`   
MiniImg       NullString       `db:"mini_img" json:"mini_img"`   
ManageEnabled       NullInt64       `db:"manage_enabled" json:"manage_enabled"`   
JobNum       NullInt64       `db:"job_num" json:"job_num"`   
Grade       NullFloat64       `db:"grade" json:"grade"`   
CustomerRange       NullString       `db:"customer_range" json:"customer_range"`   
DataRange       NullString       `db:"data_range" json:"data_range"`   
ScoreCheckEnabled       NullInt64       `db:"score_check_enabled" json:"score_check_enabled"`   
ProcessDataEnabled       NullInt64       `db:"process_data_enabled" json:"process_data_enabled"`   
ChatServiceEnabled       NullInt64       `db:"chat_service_enabled" json:"chat_service_enabled"`   
CouponConsumeEnabled       NullInt64       `db:"coupon_consume_enabled" json:"coupon_consume_enabled"`   
PositionId       NullInt64       `db:"position_id" json:"position_id"`   
Openid       NullString       `db:"openid" json:"openid"`   
CreateStaffId       NullInt64       `db:"create_staff_id" json:"create_staff_id"`   
CreateAgentId       NullInt64       `db:"create_agent_id" json:"create_agent_id"`   
}

type EStaffStore struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type EStaffStoreCust struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EStorage struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
WarehouseId       int32       `db:"warehouse_id" json:"warehouse_id"`   
LocationId       int32       `db:"location_id" json:"location_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Num       int32       `db:"num" json:"num"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EStorageRule struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
WarehouseId       int32       `db:"warehouse_id" json:"warehouse_id"`   
MinNum       int32       `db:"min_num" json:"min_num"`   
MaxNum       int32       `db:"max_num" json:"max_num"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type ESupplier struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Person       NullString       `db:"person" json:"person"`   
Phone       NullString       `db:"phone" json:"phone"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Address       NullString       `db:"address" json:"address"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
}

type ESurveyLog struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
OptometryId       int32       `db:"optometry_id" json:"optometry_id"`   
AnswerList       NullString       `db:"answer_list" json:"answer_list"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
Source       NullString       `db:"source" json:"source"`   
}

type ESurveyQuestion struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Question       NullString       `db:"question" json:"question"`   
Answer       NullString       `db:"answer" json:"answer"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type ETag struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TagName       string       `db:"tag_name" json:"tag_name"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type ETagUser struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TagId       int32       `db:"tag_id" json:"tag_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EUserActivity struct {
Id       int32       `db:"id" json:"id"`   
FromId       NullInt64       `db:"from_id" json:"from_id"`   
ToId       NullInt64       `db:"to_id" json:"to_id"`   
ActivityId       NullInt64       `db:"activity_id" json:"activity_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ChannelId       NullInt64       `db:"channel_id" json:"channel_id"`   
}

type EUserStaff struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
StaffId       int32       `db:"staff_id" json:"staff_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EUserStore struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type EUserTemplate struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
TemplateId       NullInt64       `db:"template_id" json:"template_id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
LastActiveTime       NullInt64       `db:"last_active_time" json:"last_active_time"`   
}

type EWaredocRecord struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
WarehouseDocId       int32       `db:"warehouse_doc_id" json:"warehouse_doc_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Category       NullString       `db:"category" json:"category"`   
Brand       NullString       `db:"brand" json:"brand"`   
Series       NullString       `db:"series" json:"series"`   
Model       NullString       `db:"model" json:"model"`   
Refract       NullString       `db:"refract" json:"refract"`   
Color       NullString       `db:"color" json:"color"`   
Sph       NullFloat64       `db:"sph" json:"sph"`   
Cyl       NullFloat64       `db:"cyl" json:"cyl"`   
IsAdd       NullInt64       `db:"is_add" json:"is_add"`   
Add       NullFloat64       `db:"add" json:"add"`   
Num       int32       `db:"num" json:"num"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Type       NullString       `db:"type" json:"type"`   
SubType       NullString       `db:"sub_type" json:"sub_type"`   
WarehouseId       int32       `db:"warehouse_id" json:"warehouse_id"`   
LocationId       int32       `db:"location_id" json:"location_id"`   
}

type EWarehouse struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       string       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
Manager       NullString       `db:"manager" json:"manager"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Address       NullString       `db:"address" json:"address"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
}

type EWarehouseDoc struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Sn       string       `db:"sn" json:"sn"`   
State       NullString       `db:"state" json:"state"`   
OperatorId       NullInt64       `db:"operator_id" json:"operator_id"`   
OperatorName       NullString       `db:"operator_name" json:"operator_name"`   
SendId       NullInt64       `db:"send_id" json:"send_id"`   
SendName       NullString       `db:"send_name" json:"send_name"`   
ReceiverId       NullInt64       `db:"receiver_id" json:"receiver_id"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
Type       string       `db:"type" json:"type"`   
SubType       string       `db:"sub_type" json:"sub_type"`   
FromHouseId       NullInt64       `db:"from_house_id" json:"from_house_id"`   
FromLocationId       NullInt64       `db:"from_location_id" json:"from_location_id"`   
ToHouseId       NullInt64       `db:"to_house_id" json:"to_house_id"`   
ToLocationId       NullInt64       `db:"to_location_id" json:"to_location_id"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Total       NullInt64       `db:"total" json:"total"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
ClientId       NullInt64       `db:"client_id" json:"client_id"`   
ClientName       NullString       `db:"client_name" json:"client_name"`   
SalesSn       NullString       `db:"sales_sn" json:"sales_sn"`   
}

type EWarehouseDocItems struct {
Id       int32       `db:"id" json:"id"`   
Goods       NullString       `db:"goods" json:"goods"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Num       NullInt64       `db:"num" json:"num"`   
DocId       NullInt64       `db:"doc_id" json:"doc_id"`   
Category       NullString       `db:"category" json:"category"`   
FromLocationId       NullInt64       `db:"from_location_id" json:"from_location_id"`   
FromHouseId       NullInt64       `db:"from_house_id" json:"from_house_id"`   
ToHouseId       NullInt64       `db:"to_house_id" json:"to_house_id"`   
ToLocationId       NullInt64       `db:"to_location_id" json:"to_location_id"`   
}

type EbeanRecord struct {
Id       int32       `db:"id" json:"id"`   
AuditUserRegisterId       NullInt64       `db:"audit_user_register_id" json:"audit_user_register_id"`   
BonusId       NullInt64       `db:"bonus_id" json:"bonus_id"`   
ChangeAmount       NullInt64       `db:"change_amount" json:"change_amount"`   
ChangeType       NullInt64       `db:"change_type" json:"change_type"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type Emoji struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
}

type FaceData struct {
Id       int32       `db:"id" json:"id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Third       NullString       `db:"third" json:"third"`   
Age       NullInt64       `db:"age" json:"age"`   
Gender       NullString       `db:"gender" json:"gender"`   
Raw       NullString       `db:"raw" json:"raw"`   
FileId       NullString       `db:"file_id" json:"file_id"`   
}

type FrameClientPrice struct {
Cid       int32       `db:"cid" json:"cid"`   
PriceMap       NullString       `db:"price_map" json:"price_map"`   
}

type GoodsTag struct {
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
TagId       int32       `db:"tag_id" json:"tag_id"`   
}

type GroupGoods struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Sn       string       `db:"sn" json:"sn"`   
Name       string       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
Goods       string       `db:"goods" json:"goods"`   
GroupPrice       float64       `db:"group_price" json:"group_price"`   
SumMarketPrice       float64       `db:"sum_market_price" json:"sum_market_price"`   
SumCostPrice       float64       `db:"sum_cost_price" json:"sum_cost_price"`   
SumDealerPrice       float64       `db:"sum_dealer_price" json:"sum_dealer_price"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
Price       NullFloat64       `db:"price" json:"price"`   
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
CostPrice       NullFloat64       `db:"cost_price" json:"cost_price"`   
DealerPrice       NullFloat64       `db:"dealer_price" json:"dealer_price"`   
}

type GroupLine struct {
Id       int32       `db:"id" json:"id"`   
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
GroupId       NullInt64       `db:"group_id" json:"group_id"`   
}

type LUsers struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Password       NullString       `db:"password" json:"password"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
ReceiverPhone       NullString       `db:"receiver_phone" json:"receiver_phone"`   
ReceiverAddr       NullString       `db:"receiver_addr" json:"receiver_addr"`   
ReceiverCity       NullString       `db:"receiver_city" json:"receiver_city"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
DisableCheck       NullInt64       `db:"disable_check" json:"disable_check"`   
IsDealer       NullInt64       `db:"is_dealer" json:"is_dealer"`   
BillPassword       NullString       `db:"bill_password" json:"bill_password"`   
}

type Len struct {
Id       int32       `db:"id" json:"id"`   
Refract       NullFloat64       `db:"refract" json:"refract"`   
Brand       string       `db:"brand" json:"brand"`   
Category       string       `db:"category" json:"category"`   
LumLow       NullInt64       `db:"lum_low" json:"lum_low"`   
LumHigh       NullInt64       `db:"lum_high" json:"lum_high"`   
Price       NullInt64       `db:"price" json:"price"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
ShopOnUse       int32       `db:"shop_on_use" json:"shop_on_use"`   
OnSale       int32       `db:"on_sale" json:"on_sale"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
Sales       NullInt64       `db:"sales" json:"sales"`   
Loves       NullInt64       `db:"loves" json:"loves"`   
}

type LenBrand struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
Img4       NullString       `db:"img_4" json:"img_4"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ThirdId       int32       `db:"third_id" json:"third_id"`   
Refract       float64       `db:"refract" json:"refract"`   
}

type LenCategory struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
DetailFigure1       NullString       `db:"detail_figure_1" json:"detail_figure_1"`   
DetailFigure2       NullString       `db:"detail_figure_2" json:"detail_figure_2"`   
DetailFigure3       NullString       `db:"detail_figure_3" json:"detail_figure_3"`   
DetailFigure4       NullString       `db:"detail_figure_4" json:"detail_figure_4"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Refract       float64       `db:"refract" json:"refract"`   
Category       string       `db:"category" json:"category"`   
Name       NullString       `db:"name" json:"name"`   
Thumb       NullString       `db:"thumb" json:"thumb"`   
HasPackage       NullInt64       `db:"has_package" json:"has_package"`   
PackagePrice       NullFloat64       `db:"package_price" json:"package_price"`   
}

type Lens struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
Category       string       `db:"category" json:"category"`   
Refract       float64       `db:"refract" json:"refract"`   
Sn       string       `db:"sn" json:"sn"`   
Color       string       `db:"color" json:"color"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDelete       NullInt64       `db:"is_delete" json:"is_delete"`   
DealerPrice       float64       `db:"dealer_price" json:"dealer_price"`   
CostPrice       NullFloat64       `db:"cost_price" json:"cost_price"`   
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
ShopName       NullString       `db:"shop_name" json:"shop_name"`   
ColorName       NullString       `db:"color_name" json:"color_name"`   
Num       int32       `db:"num" json:"num"`   
SalesNum       int32       `db:"sales_num" json:"sales_num"`   
LoveNum       int32       `db:"love_num" json:"love_num"`   
Factory       NullString       `db:"factory" json:"factory"`   
SupplierId       NullInt64       `db:"supplier_id" json:"supplier_id"`   
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
}

type LensConvexConcave struct {
Id       int32       `db:"id" json:"id"`   
LensId       int32       `db:"lens_id" json:"lens_id"`   
BallLens       NullFloat64       `db:"ball_lens" json:"ball_lens"`   
PillarLens       NullFloat64       `db:"pillar_lens" json:"pillar_lens"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDelete       NullInt64       `db:"is_delete" json:"is_delete"`   
Stock       NullInt64       `db:"stock" json:"stock"`   
}

type LogoPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MCenter struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Password       NullString       `db:"password" json:"password"`   
CenterName       NullString       `db:"center_name" json:"center_name"`   
ExpireTime       NullTime       `db:"expire_time" json:"expire_time"`   // can be null, use with caution
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Granted       NullInt64       `db:"granted" json:"granted"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Street       NullString       `db:"street" json:"street"`   
Address       NullString       `db:"address" json:"address"`   
Coordinate       NullString       `db:"coordinate" json:"coordinate"`   
ShopNo       NullInt64       `db:"shop_no" json:"shop_no"`   
CityCode       NullInt64       `db:"city_code" json:"city_code"`   
ContactName       NullString       `db:"contact_name" json:"contact_name"`   
ContactPhone       NullString       `db:"contact_phone" json:"contact_phone"`   
UrgentFee       NullFloat64       `db:"urgent_fee" json:"urgent_fee"`   
CommonFee       NullFloat64       `db:"common_fee" json:"common_fee"`   
UrgentFeeInfo       NullString       `db:"urgent_fee_info" json:"urgent_fee_info"`   
}

type MCenterFee struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CenterId       int32       `db:"center_id" json:"center_id"`   
Price       NullFloat64       `db:"price" json:"price"`   
Rank       NullInt64       `db:"rank" json:"rank"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type MCenterInsurance struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CenterId       int32       `db:"center_id" json:"center_id"`   
Price       NullFloat64       `db:"price" json:"price"`   
Rank       NullInt64       `db:"rank" json:"rank"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type MCenterStore struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CenterId       int32       `db:"center_id" json:"center_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
UrgentFee       NullFloat64       `db:"urgent_fee" json:"urgent_fee"`   
CommonFee       NullFloat64       `db:"common_fee" json:"common_fee"`   
ShippingFee       NullFloat64       `db:"shipping_fee" json:"shipping_fee"`   
}

type MDeliver struct {
Id       int32       `db:"id" json:"id"`   
ClientSn       NullString       `db:"client_sn" json:"client_sn"`   
OriginId       NullString       `db:"origin_id" json:"origin_id"`   
Provider       NullString       `db:"provider" json:"provider"`   
Direction       string       `db:"direction" json:"direction"`   
CenterId       NullInt64       `db:"center_id" json:"center_id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
ShippingFee       NullFloat64       `db:"shipping_fee" json:"shipping_fee"`   
TransportName       NullString       `db:"transport_name" json:"transport_name"`   
TransportAddress       NullString       `db:"transport_address" json:"transport_address"`   
TransportLongitude       NullInt64       `db:"transport_longitude" json:"transport_longitude"`   
TransportLatitude       NullInt64       `db:"transport_latitude" json:"transport_latitude"`   
TransportTel       NullString       `db:"transport_tel" json:"transport_tel"`   
TransportRemark       NullString       `db:"transport_remark" json:"transport_remark"`   
CargoPrice       NullFloat64       `db:"cargo_price" json:"cargo_price"`   
IsInvoiced       NullString       `db:"is_invoiced" json:"is_invoiced"`   
OrderPaymentStatus       NullInt64       `db:"order_payment_status" json:"order_payment_status"`   
OrderPaymentMethod       NullInt64       `db:"order_payment_method" json:"order_payment_method"`   
IsAgentPayment       NullInt64       `db:"is_agent_payment" json:"is_agent_payment"`   
GoodsCount       NullInt64       `db:"goods_count" json:"goods_count"`   
RequireReceiveTime       NullInt64       `db:"require_receive_time" json:"require_receive_time"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
ReceiverPhone       NullString       `db:"receiver_phone" json:"receiver_phone"`   
ReceiverAddress       NullString       `db:"receiver_address" json:"receiver_address"`   
ReceiverLng       NullInt64       `db:"receiver_lng" json:"receiver_lng"`   
ReceiverLat       NullInt64       `db:"receiver_lat" json:"receiver_lat"`   
OrderStatus       NullInt64       `db:"order_status" json:"order_status"`   
StatusTime       NullTime       `db:"status_time" json:"status_time"`   // can be null, use with caution
DriverId       NullInt64       `db:"driver_id" json:"driver_id"`   
DriverName       NullString       `db:"driver_name" json:"driver_name"`   
DriverPhone       NullString       `db:"driver_phone" json:"driver_phone"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
Creator       NullString       `db:"creator" json:"creator"`   
CreatorId       NullInt64       `db:"creator_id" json:"creator_id"`   
CancelFrom       NullString       `db:"cancel_from" json:"cancel_from"`   
CancelReason       NullString       `db:"cancel_reason" json:"cancel_reason"`   
}

type MDeliverState struct {
Id       int32       `db:"id" json:"id"`   
DeliverId       int32       `db:"deliver_id" json:"deliver_id"`   
StatusCode       string       `db:"status_code" json:"status_code"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
StatusTime       NullTime       `db:"status_time" json:"status_time"`   // can be null, use with caution
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MGoods struct {
Id       int32       `db:"id" json:"id"`   
Category       NullString       `db:"category" json:"category"`   
Brand       string       `db:"brand" json:"brand"`   
Serial       string       `db:"serial" json:"serial"`   
Name       string       `db:"name" json:"name"`   
IsCustomize       NullInt64       `db:"is_customize" json:"is_customize"`   
Refraction       NullFloat64       `db:"refraction" json:"refraction"`   
Material       NullString       `db:"material" json:"material"`   
Surface       NullInt64       `db:"surface" json:"surface"`   
Sph       NullFloat64       `db:"sph" json:"sph"`   
SphMin       NullFloat64       `db:"sph_min" json:"sph_min"`   
SphMax       NullFloat64       `db:"sph_max" json:"sph_max"`   
Cyl       NullFloat64       `db:"cyl" json:"cyl"`   
CylMin       NullFloat64       `db:"cyl_min" json:"cyl_min"`   
CylMax       NullFloat64       `db:"cyl_max" json:"cyl_max"`   
LensAdd       NullFloat64       `db:"lens_add" json:"lens_add"`   
AddMin       NullFloat64       `db:"add_min" json:"add_min"`   
AddMax       NullFloat64       `db:"add_max" json:"add_max"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MGoodsCenter struct {
Id       int32       `db:"id" json:"id"`   
CenterId       int32       `db:"center_id" json:"center_id"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
TradePrice       float64       `db:"trade_price" json:"trade_price"`   
CostPrice       float64       `db:"cost_price" json:"cost_price"`   
GoodsStatus       int32       `db:"goods_status" json:"goods_status"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
}

type MPayLog struct {
Id       int32       `db:"id" json:"id"`   
PaySn       NullString       `db:"pay_sn" json:"pay_sn"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
CenterId       NullInt64       `db:"center_id" json:"center_id"`   
Event       NullString       `db:"event" json:"event"`   
TradeAccount       NullString       `db:"trade_account" json:"trade_account"`   
Money       NullFloat64       `db:"money" json:"money"`   
DeductFee       NullFloat64       `db:"deduct_fee" json:"deduct_fee"`   
ProcessId       NullInt64       `db:"process_id" json:"process_id"`   
DeliverId       NullInt64       `db:"deliver_id" json:"deliver_id"`   
Price       NullFloat64       `db:"price" json:"price"`   
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
StoreBalance       NullFloat64       `db:"store_balance" json:"store_balance"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MPayState struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
GoodsType       string       `db:"goods_type" json:"goods_type"`   
GoodsDetail       string       `db:"goods_detail" json:"goods_detail"`   
Money       float64       `db:"money" json:"money"`   
Status       int32       `db:"status" json:"status"`   
TradeNo       string       `db:"trade_no" json:"trade_no"`   
Hash       string       `db:"hash" json:"hash"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
MProcessIds       NullString       `db:"m_process_ids" json:"m_process_ids"`   
Source       NullString       `db:"source" json:"source"`   
CenterId       NullInt64       `db:"center_id" json:"center_id"`   
CreatorId       NullInt64       `db:"creator_id" json:"creator_id"`   
}

type MProcess struct {
Id       int32       `db:"id" json:"id"`   
Sn       string       `db:"sn" json:"sn"`   
ShortSn       string       `db:"short_sn" json:"short_sn"`   
CenterId       NullInt64       `db:"center_id" json:"center_id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
Remark       NullString       `db:"remark" json:"remark"`   
Status       NullString       `db:"status" json:"status"`   
GoodsPrice       NullFloat64       `db:"goods_price" json:"goods_price"`   
ProcessFee       NullFloat64       `db:"process_fee" json:"process_fee"`   
FinalPrice       NullFloat64       `db:"final_price" json:"final_price"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
UserPhone       NullString       `db:"user_phone" json:"user_phone"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
Relation       NullString       `db:"relation" json:"relation"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
StoreName       NullString       `db:"store_name" json:"store_name"`   
StoreAddress       NullString       `db:"store_address" json:"store_address"`   
StoreCoordinate       NullString       `db:"store_coordinate" json:"store_coordinate"`   
StoreTelPhone       NullString       `db:"store_tel_phone" json:"store_tel_phone"`   
StoreStaffId       NullInt64       `db:"store_staff_id" json:"store_staff_id"`   
StoreStaffName       NullString       `db:"store_staff_name" json:"store_staff_name"`   
OptStaffId       NullInt64       `db:"opt_staff_id" json:"opt_staff_id"`   
InsptStaffId       NullInt64       `db:"inspt_staff_id" json:"inspt_staff_id"`   
OptStaffName       NullString       `db:"opt_staff_name" json:"opt_staff_name"`   
InsptStaffName       NullString       `db:"inspt_staff_name" json:"inspt_staff_name"`   
OptometryId       NullInt64       `db:"optometry_id" json:"optometry_id"`   
OptometrySn       NullString       `db:"optometry_sn" json:"optometry_sn"`   
OptometryTime       NullTime       `db:"optometry_time" json:"optometry_time"`   // can be null, use with caution
SphLeft       NullFloat64       `db:"sph_left" json:"sph_left"`   
SphRight       NullFloat64       `db:"sph_right" json:"sph_right"`   
CylLeft       NullFloat64       `db:"cyl_left" json:"cyl_left"`   
CylRight       NullFloat64       `db:"cyl_right" json:"cyl_right"`   
AxisRight       NullInt64       `db:"axis_right" json:"axis_right"`   
AxisLeft       NullInt64       `db:"axis_left" json:"axis_left"`   
Pd       NullFloat64       `db:"pd" json:"pd"`   
PdLeft       NullFloat64       `db:"pd_left" json:"pd_left"`   
PdRight       NullFloat64       `db:"pd_right" json:"pd_right"`   
Ph       NullInt64       `db:"ph" json:"ph"`   
PhLeft       NullInt64       `db:"ph_left" json:"ph_left"`   
PhRight       NullInt64       `db:"ph_right" json:"ph_right"`   
PrismLeft       NullInt64       `db:"prism_left" json:"prism_left"`   
PrismRight       NullInt64       `db:"prism_right" json:"prism_right"`   
AddLeft       NullFloat64       `db:"add_left" json:"add_left"`   
AddRight       NullFloat64       `db:"add_right" json:"add_right"`   
StatusLeft       NullString       `db:"status_left" json:"status_left"`   
StatusRight       NullString       `db:"status_right" json:"status_right"`   
DomainEye       NullString       `db:"domain_eye" json:"domain_eye"`   
IsSame       NullInt64       `db:"is_same" json:"is_same"`   
CorctVisionLeft       NullFloat64       `db:"corct_vision_left" json:"corct_vision_left"`   
CorctVisionRight       NullFloat64       `db:"corct_vision_right" json:"corct_vision_right"`   
OriginVisionLeft       NullFloat64       `db:"origin_vision_left" json:"origin_vision_left"`   
OriginVisionRight       NullFloat64       `db:"origin_vision_right" json:"origin_vision_right"`   
ConsumptionId       NullInt64       `db:"consumption_id" json:"consumption_id"`   
ConsumptionSn       NullString       `db:"consumption_sn" json:"consumption_sn"`   
GoodsInfo       NullString       `db:"goods_info" json:"goods_info"`   
ConsumptionTime       NullTime       `db:"consumption_time" json:"consumption_time"`   // can be null, use with caution
IsUrgent       NullInt64       `db:"is_urgent" json:"is_urgent"`   
DeliverTime       NullTime       `db:"deliver_time" json:"deliver_time"`   // can be null, use with caution
Left       NullString       `db:"left" json:"left"`   
Right       NullString       `db:"right" json:"right"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
IsNotifying       NullInt64       `db:"is_notifying" json:"is_notifying"`   
FrameImg       NullString       `db:"frame_img" json:"frame_img"`   
OrderPlace       NullString       `db:"order_place" json:"order_place"`   
CreateStaffId       NullInt64       `db:"create_staff_id" json:"create_staff_id"`   
CreateStaffName       NullString       `db:"create_staff_name" json:"create_staff_name"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CancelReason       NullString       `db:"cancel_reason" json:"cancel_reason"`   
CancelFrom       NullString       `db:"cancel_from" json:"cancel_from"`   
OptometryImg       NullString       `db:"optometry_img" json:"optometry_img"`   
Source       NullString       `db:"source" json:"source"`   
LeftInfo       NullString       `db:"left_info" json:"left_info"`   
RightInfo       NullString       `db:"right_info" json:"right_info"`   
ProcessFeeDesc       NullString       `db:"process_fee_desc" json:"process_fee_desc"`   
InsuranceId       NullInt64       `db:"insurance_id" json:"insurance_id"`   
InsuranceFee       NullFloat64       `db:"insurance_fee" json:"insurance_fee"`   
InsuranceFeeName       NullString       `db:"insurance_fee_name" json:"insurance_fee_name"`   
InsuranceFeeDesc       NullString       `db:"insurance_fee_desc" json:"insurance_fee_desc"`   
ProcessFeeId       NullInt64       `db:"process_fee_id" json:"process_fee_id"`   
UrgentFee       NullFloat64       `db:"urgent_fee" json:"urgent_fee"`   
UrgentFeeInfo       NullString       `db:"urgent_fee_info" json:"urgent_fee_info"`   
ProcessFeeName       NullString       `db:"process_fee_name" json:"process_fee_name"`   
}

type MProcessDeliver struct {
Id       int32       `db:"id" json:"id"`   
CenterId       int32       `db:"center_id" json:"center_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
StoreId       int32       `db:"store_id" json:"store_id"`   
ProcessId       int32       `db:"process_id" json:"process_id"`   
DeliverId       int32       `db:"deliver_id" json:"deliver_id"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MProcessGoods struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
ProcessId       NullInt64       `db:"process_id" json:"process_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
Brand       NullString       `db:"brand" json:"brand"`   
GoodsInfo       NullString       `db:"goods_info" json:"goods_info"`   
Refraction       NullFloat64       `db:"refraction" json:"refraction"`   
Sph       NullFloat64       `db:"sph" json:"sph"`   
Cyl       NullFloat64       `db:"cyl" json:"cyl"`   
Add       NullFloat64       `db:"add" json:"add"`   
IsCustomize       NullString       `db:"is_customize" json:"is_customize"`   
Num       NullInt64       `db:"num" json:"num"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Price       NullFloat64       `db:"price" json:"price"`   
IsLeft       NullInt64       `db:"is_left" json:"is_left"`   
IsRight       NullInt64       `db:"is_right" json:"is_right"`   
}

type MProcessStatus struct {
Id       int32       `db:"id" json:"id"`   
ProcessId       NullInt64       `db:"process_id" json:"process_id"`   
OperatorId       NullInt64       `db:"operator_id" json:"operator_id"`   
OperatorName       NullString       `db:"operator_name" json:"operator_name"`   
Status       NullString       `db:"status" json:"status"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type MStaff struct {
Id       int32       `db:"id" json:"id"`   
CenterId       int32       `db:"center_id" json:"center_id"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
Account       NullString       `db:"account" json:"account"`   
Password       NullString       `db:"password" json:"password"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Position       NullInt64       `db:"position" json:"position"`   
OperateEnabled       NullInt64       `db:"operate_enabled" json:"operate_enabled"`   
InspectEnabled       NullInt64       `db:"inspect_enabled" json:"inspect_enabled"`   
ShippingEnabled       NullInt64       `db:"shipping_enabled" json:"shipping_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Openid       NullString       `db:"openid" json:"openid"`   
}

type ModelPos struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Width       NullInt64       `db:"width" json:"width"`   
Pos       NullString       `db:"pos" json:"pos"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NAccountant struct {
Id       int32       `db:"id" json:"id"`   
Name       string       `db:"name" json:"name"`   
Phone       string       `db:"phone" json:"phone"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Password       NullString       `db:"password" json:"password"`   
}

type NAdmin struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Password       NullString       `db:"password" json:"password"`   
Name       NullString       `db:"name" json:"name"`   
Role       NullString       `db:"role" json:"role"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Granted       NullInt64       `db:"granted" json:"granted"`   
ThirdEnabled       NullString       `db:"third_enabled" json:"third_enabled"`   
CourseEnabled       NullString       `db:"course_enabled" json:"course_enabled"`   
ScoreGoodsEnabled       NullString       `db:"score_goods_enabled" json:"score_goods_enabled"`   
ScoreMiniProgramEnabled       NullString       `db:"score_mini_program_enabled" json:"score_mini_program_enabled"`   
ProcessEnabled       NullString       `db:"process_enabled" json:"process_enabled"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
AnnounceEnabled       NullString       `db:"announce_enabled" json:"announce_enabled"`   
}

type NAuditor struct {
WxId       string       `db:"wx_id" json:"wx_id"`   
Name       NullString       `db:"name" json:"name"`   
}

type NBonusRecord struct {
Id       int32       `db:"id" json:"id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
TotalMoney       NullInt64       `db:"total_money" json:"total_money"`   
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
IsRandom       NullInt64       `db:"is_random" json:"is_random"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
IsActivate       NullInt64       `db:"is_activate" json:"is_activate"`   
NeedApproval       NullInt64       `db:"need_approval" json:"need_approval"`   
}

type NFaceRaw struct {
Id       int32       `db:"id" json:"id"`   
FileId       NullString       `db:"file_id" json:"file_id"`   
Raw       NullString       `db:"raw" json:"raw"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NFaceResult struct {
Id       int32       `db:"id" json:"id"`   
FileId       NullString       `db:"file_id" json:"file_id"`   
Result       NullString       `db:"result" json:"result"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Src       NullString       `db:"src" json:"src"`   
}

type NFirstPrice struct {
Id       int32       `db:"id" json:"id"`   
Price       NullFloat64       `db:"price" json:"price"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
ClientId       NullInt64       `db:"client_id" json:"client_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NFrameCharge struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Money       int32       `db:"money" json:"money"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ClientId       int32       `db:"client_id" json:"client_id"`   
ClientName       string       `db:"client_name" json:"client_name"`   
Sn       NullString       `db:"sn" json:"sn"`   
Type       NullString       `db:"type" json:"type"`   
}

type NFrameInputLine struct {
Id       int32       `db:"id" json:"id"`   
Code       NullString       `db:"code" json:"code"`   
Sn       NullString       `db:"sn" json:"sn"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
Operator       NullString       `db:"operator" json:"operator"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
Factory       NullString       `db:"factory" json:"factory"`   
Brand       NullString       `db:"brand" json:"brand"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
Status       NullString       `db:"status" json:"status"`   
IsLens       NullInt64       `db:"is_lens" json:"is_lens"`   
}

type NFrameOrder struct {
Id       int32       `db:"id" json:"id"`   
Sn       NullString       `db:"sn" json:"sn"`   
TotalPrice       float64       `db:"total_price" json:"total_price"`   
ClientId       int32       `db:"client_id" json:"client_id"`   
User       string       `db:"user" json:"user"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
UserStore       NullString       `db:"user_store" json:"user_store"`   
UserAddr       NullString       `db:"user_addr" json:"user_addr"`   
Third       NullString       `db:"third" json:"third"`   
ThirdName       NullString       `db:"third_name" json:"third_name"`   
Status       NullString       `db:"status" json:"status"`   
TotalGlasses       int32       `db:"total_glasses" json:"total_glasses"`   
PayMethod       NullString       `db:"pay_method" json:"pay_method"`   
ShipCompany       NullString       `db:"ship_company" json:"ship_company"`   
ShipNo       NullString       `db:"ship_no" json:"ship_no"`   
ShipRemark       NullString       `db:"ship_remark" json:"ship_remark"`   
ShipFee       NullFloat64       `db:"ship_fee" json:"ship_fee"`   
ShipCreateTime       NullString       `db:"ship_create_time" json:"ship_create_time"`   
ShipUpdateTime       NullTime       `db:"ship_update_time" json:"ship_update_time"`   // can be null, use with caution
CreateTime       NullString       `db:"create_time" json:"create_time"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
Remark       NullString       `db:"remark" json:"remark"`   
StaffPhone       string       `db:"staff_phone" json:"staff_phone"`   
StaffName       NullString       `db:"staff_name" json:"staff_name"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
Urgency       NullString       `db:"urgency" json:"urgency"`   
AddRemark       NullString       `db:"add_remark" json:"add_remark"`   
DiscountPrice       NullFloat64       `db:"discount_price" json:"discount_price"`   
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
GroupData       NullString       `db:"group_data" json:"group_data"`   
FinalTotalPrice       NullFloat64       `db:"final_total_price" json:"final_total_price"`   
}

type NFrameOrderLine struct {
Id       int32       `db:"id" json:"id"`   
OrderId       int32       `db:"order_id" json:"order_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
Num       int32       `db:"num" json:"num"`   
GoodsImg       NullString       `db:"goods_img" json:"goods_img"`   
CreateTime       NullString       `db:"create_time" json:"create_time"`   
Name       string       `db:"name" json:"name"`   
OriginalPrice       int32       `db:"original_price" json:"original_price"`   
TotalPrice       float64       `db:"total_price" json:"total_price"`   
Color       NullString       `db:"color" json:"color"`   
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
Status       NullString       `db:"status" json:"status"`   
StoreOrderId       NullInt64       `db:"store_order_id" json:"store_order_id"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Sn       NullString       `db:"sn" json:"sn"`   
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
IsAdd       NullInt64       `db:"is_add" json:"is_add"`   
IsNew       NullInt64       `db:"is_new" json:"is_new"`   
Finished       NullInt64       `db:"finished" json:"finished"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
LackStock       NullInt64       `db:"lackStock" json:"lackStock"`   
CodeIds       NullString       `db:"code_ids" json:"code_ids"`   
IsLens       NullInt64       `db:"is_lens" json:"is_lens"`   
IsGroup       int32       `db:"is_group" json:"is_group"`   
Optometries       NullString       `db:"optometries" json:"optometries"`   
Refract       NullFloat64       `db:"refract" json:"refract"`   
LensId       NullInt64       `db:"lens_id" json:"lens_id"`   
BpId       NullInt64       `db:"bp_id" json:"bp_id"`   
GroupId       NullInt64       `db:"group_id" json:"group_id"`   
Remark       NullString       `db:"remark" json:"remark"`   
Optometry       NullString       `db:"optometry" json:"optometry"`   
}

type NFrameOrderReturn struct {
Id       int32       `db:"id" json:"id"`   
OrderId       NullInt64       `db:"order_id" json:"order_id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
UploadImg       NullString       `db:"upload_img" json:"upload_img"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NFrameReturnLine struct {
Id       int32       `db:"id" json:"id"`   
OrderId       NullInt64       `db:"order_id" json:"order_id"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Price       NullFloat64       `db:"price" json:"price"`   
Num       NullInt64       `db:"num" json:"num"`   
TotalPrice       NullFloat64       `db:"total_price" json:"total_price"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NGoods2 struct {
Id       int32       `db:"id" json:"id"`   
Sn       NullString       `db:"sn" json:"sn"`   
OnSale       NullInt64       `db:"on_sale" json:"on_sale"`   
ShopPrice       NullFloat64       `db:"shop_price" json:"shop_price"`   
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
Name       NullString       `db:"name" json:"name"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
ImgFace       NullString       `db:"img_face" json:"img_face"`   
ImgLeg       NullString       `db:"img_leg" json:"img_leg"`   
Description       NullString       `db:"description" json:"description"`   
Category       NullString       `db:"category" json:"category"`   
Color       NullString       `db:"color" json:"color"`   
Style       NullString       `db:"style" json:"style"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Zone       NullString       `db:"zone" json:"zone"`   
Recommend       NullInt64       `db:"recommend" json:"recommend"`   
Saled       NullInt64       `db:"saled" json:"saled"`   
Stock       NullInt64       `db:"stock" json:"stock"`   
Brand       NullString       `db:"brand" json:"brand"`   
LinkUrl       NullString       `db:"link_url" json:"link_url"`   
TryOwner       NullString       `db:"try_owner" json:"try_owner"`   
TryDesc       NullString       `db:"try_desc" json:"try_desc"`   
ThirdStock       NullInt64       `db:"third_stock" json:"third_stock"`   
Attrs       NullString       `db:"attrs" json:"attrs"`   
PriceRegion       NullString       `db:"price_region" json:"price_region"`   
Frame       NullString       `db:"frame" json:"frame"`   
FrameStyle       NullString       `db:"frame_style" json:"frame_style"`   
Material       NullString       `db:"material" json:"material"`   
FaceStyle       NullString       `db:"face_style" json:"face_style"`   
Price       NullFloat64       `db:"price" json:"price"`   
DealerPrice       NullFloat64       `db:"dealer_price" json:"dealer_price"`   
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
Factory       NullString       `db:"factory" json:"factory"`   
TryOwnerId       NullInt64       `db:"try_owner_id" json:"try_owner_id"`   
GlassType       NullString       `db:"glass_type" json:"glass_type"`   
Sex       NullString       `db:"sex" json:"sex"`   
CostPrice       NullFloat64       `db:"cost_price" json:"cost_price"`   
SupplierId       NullInt64       `db:"supplier_id" json:"supplier_id"`   
Refract       NullFloat64       `db:"refract" json:"refract"`   
IsLens       NullInt64       `db:"is_lens" json:"is_lens"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
SourceId       NullInt64       `db:"source_id" json:"source_id"`   
}

type NGoodsColorImages struct {
Id       int32       `db:"id" json:"id"`   
Color       string       `db:"color" json:"color"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Sn       NullString       `db:"sn" json:"sn"`   
ImgDetail0       NullString       `db:"img_detail_0" json:"img_detail_0"`   
ImgDetail1       NullString       `db:"img_detail_1" json:"img_detail_1"`   
ImgDetail2       NullString       `db:"img_detail_2" json:"img_detail_2"`   
ImgDetail3       NullString       `db:"img_detail_3" json:"img_detail_3"`   
ImgDetail4       NullString       `db:"img_detail_4" json:"img_detail_4"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
ImgLeg       NullString       `db:"img_leg" json:"img_leg"`   
ImgFace       NullString       `db:"img_face" json:"img_face"`   
Remark       NullString       `db:"remark" json:"remark"`   
LegOrigin       NullString       `db:"leg_origin" json:"leg_origin"`   
FaceOrigin       NullString       `db:"face_origin" json:"face_origin"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Stock       NullInt64       `db:"stock" json:"stock"`   
Price       NullFloat64       `db:"price" json:"price"`   
UpperLeftX       NullInt64       `db:"upper_left_x" json:"upper_left_x"`   
UpperLeftY       NullInt64       `db:"upper_left_y" json:"upper_left_y"`   
UpperRightX       NullInt64       `db:"upper_right_x" json:"upper_right_x"`   
UpperRightY       NullInt64       `db:"upper_right_y" json:"upper_right_y"`   
LowerLeftX       NullInt64       `db:"lower_left_x" json:"lower_left_x"`   
LowerLeftY       NullInt64       `db:"lower_left_y" json:"lower_left_y"`   
LowerRightX       NullInt64       `db:"lower_right_x" json:"lower_right_x"`   
LowerRightY       NullInt64       `db:"lower_right_y" json:"lower_right_y"`   
LegY1       NullInt64       `db:"leg_y1" json:"leg_y1"`   
LegY2       NullInt64       `db:"leg_y2" json:"leg_y2"`   
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
Factory       NullString       `db:"factory" json:"factory"`   
TryOnUse       NullInt64       `db:"try_on_use" json:"try_on_use"`   
BonusOnUse       NullInt64       `db:"bonus_on_use" json:"bonus_on_use"`   
IsNew       NullInt64       `db:"is_new" json:"is_new"`   
ShopOnUse       NullInt64       `db:"shop_on_use" json:"shop_on_use"`   
Sales       NullInt64       `db:"sales" json:"sales"`   
Loves       NullInt64       `db:"loves" json:"loves"`   
OnSale       NullInt64       `db:"on_sale" json:"on_sale"`   
Name       NullString       `db:"name" json:"name"`   
TryOnline       NullInt64       `db:"try_online" json:"try_online"`   
Invisible       NullInt64       `db:"invisible" json:"invisible"`   
DiscountPrice       NullFloat64       `db:"discount_price" json:"discount_price"`   
BallLens       NullFloat64       `db:"ball_lens" json:"ball_lens"`   
PillarLens       NullFloat64       `db:"pillar_lens" json:"pillar_lens"`   
IsDeleted       int32       `db:"is_deleted" json:"is_deleted"`   
BuyDetail1       NullString       `db:"buy_detail_1" json:"buy_detail_1"`   
BuyDetail2       NullString       `db:"buy_detail_2" json:"buy_detail_2"`   
BuyDetail3       NullString       `db:"buy_detail_3" json:"buy_detail_3"`   
BuyDetail4       NullString       `db:"buy_detail_4" json:"buy_detail_4"`   
SourceId       NullInt64       `db:"source_id" json:"source_id"`   
}

type NGoodsLens struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
Pic       NullString       `db:"pic" json:"pic"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Data       NullString       `db:"data" json:"data"`   
Attr1       NullString       `db:"attr1" json:"attr1"`   
Value1       NullString       `db:"value1" json:"value1"`   
MultiAttrs       NullString       `db:"multi_attrs" json:"multi_attrs"`   
Sold       NullInt64       `db:"sold" json:"sold"`   
Customs       NullString       `db:"customs" json:"customs"`   
MaServices       NullString       `db:"ma_services" json:"ma_services"`   
}

type NGoodsLensCustom struct {
Id       int32       `db:"id" json:"id"`   
User       NullString       `db:"user" json:"user"`   
LensId       NullInt64       `db:"lens_id" json:"lens_id"`   
Name       NullString       `db:"name" json:"name"`   
Custom       NullString       `db:"custom" json:"custom"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CustomServices       NullString       `db:"custom_services" json:"custom_services"`   
}

type NLens struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Price       NullFloat64       `db:"price" json:"price"`   
Description       NullString       `db:"description" json:"description"`   
}

type NLensAttrs struct {
Id       int32       `db:"id" json:"id"`   
Attrs       NullString       `db:"attrs" json:"attrs"`   
}

type NLensCat struct {
Id       int32       `db:"id" json:"id"`   
Type       string       `db:"type" json:"type"`   
Fraction       string       `db:"fraction" json:"fraction"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NLensLineOrder struct {
Id       int32       `db:"id" json:"id"`   
LensOrderId       int32       `db:"lens_order_id" json:"lens_order_id"`   
LenId       int32       `db:"len_id" json:"len_id"`   
Num       int32       `db:"num" json:"num"`   
GoodsImg       NullString       `db:"goods_img" json:"goods_img"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Eye       NullString       `db:"eye" json:"eye"`   
Price       float64       `db:"price" json:"price"`   
Remark       NullString       `db:"remark" json:"remark"`   
Refract       NullString       `db:"refract" json:"refract"`   
Sph       NullString       `db:"sph" json:"sph"`   
Cyl       NullString       `db:"cyl" json:"cyl"`   
Pic       NullString       `db:"pic" json:"pic"`   
Name       NullString       `db:"name" json:"name"`   
Attrs       NullString       `db:"attrs" json:"attrs"`   
Axis       NullString       `db:"axis" json:"axis"`   
ShipNo       NullString       `db:"ship_no" json:"ship_no"`   
ShipDate       NullTime       `db:"ship_date" json:"ship_date"`   // can be null, use with caution
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
Status       NullString       `db:"status" json:"status"`   
Services       NullString       `db:"services" json:"services"`   
ShipFee       NullInt64       `db:"ship_fee" json:"ship_fee"`   
ShipNotice       NullString       `db:"ship_notice" json:"ship_notice"`   
ReturnOrderId       NullInt64       `db:"return_order_id" json:"return_order_id"`   
}

type NLensOrder struct {
Id       int32       `db:"id" json:"id"`   
Sn       string       `db:"sn" json:"sn"`   
TotalPrice       float64       `db:"total_price" json:"total_price"`   
UserPhone       string       `db:"user_phone" json:"user_phone"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
Status       NullString       `db:"status" json:"status"`   
Paid       NullInt64       `db:"paid" json:"paid"`   
ShipStatus       NullString       `db:"ship_status" json:"ship_status"`   
ShipNo       NullString       `db:"ship_no" json:"ship_no"`   
ShipFee       NullFloat64       `db:"ship_fee" json:"ship_fee"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Remark       NullString       `db:"remark" json:"remark"`   
BuyReceiver       NullString       `db:"buy_receiver" json:"buy_receiver"`   
BuyReceiverPhone       NullString       `db:"buy_receiver_phone" json:"buy_receiver_phone"`   
BuyAddress       NullString       `db:"buy_address" json:"buy_address"`   
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
ShipProv       NullString       `db:"ship_prov" json:"ship_prov"`   
PayQuery       NullInt64       `db:"pay_query" json:"pay_query"`   
Source       NullString       `db:"source" json:"source"`   
}

type NLensPrice struct {
Id       int32       `db:"id" json:"id"`   
Type       NullString       `db:"type" json:"type"`   
Brand       NullString       `db:"brand" json:"brand"`   
Price       NullFloat64       `db:"price" json:"price"`   
Order1       NullInt64       `db:"order1" json:"order1"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
MarketPrice       NullFloat64       `db:"market_price" json:"market_price"`   
}

type NOpenApp struct {
Id       int32       `db:"id" json:"id"`   
OpenAppid       string       `db:"open_appid" json:"open_appid"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
WxAppid       string       `db:"wx_appid" json:"wx_appid"`   
MicroShopAppid       NullString       `db:"micro_shop_appid" json:"micro_shop_appid"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CodeDesc       NullString       `db:"code_desc" json:"code_desc"`   
}

type NOrderRedpack struct {
Id       int32       `db:"id" json:"id"`   
UserPhone       NullString       `db:"user_phone" json:"user_phone"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
ChargePrice       NullFloat64       `db:"charge_price" json:"charge_price"`   
Sn       NullString       `db:"sn" json:"sn"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Paid       NullInt64       `db:"paid" json:"paid"`   
PayQuery       NullInt64       `db:"pay_query" json:"pay_query"`   
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
}

type NPlatformWechatTemplate struct {
Id       int32       `db:"id" json:"id"`   
Appid       NullString       `db:"appid" json:"appid"`   
ShortId       NullString       `db:"short_id" json:"short_id"`   
TemplateId       NullString       `db:"template_id" json:"template_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NQueue2 struct {
Id       int32       `db:"id" json:"id"`   
User       NullString       `db:"user" json:"user"`   
QueueName       NullString       `db:"queue_name" json:"queue_name"`   
Status       NullString       `db:"status" json:"status"`   
UserPhone       NullString       `db:"user_phone" json:"user_phone"`   
Notified       NullInt64       `db:"notified" json:"notified"`   
Rank       NullInt64       `db:"rank" json:"rank"`   
Expert       NullString       `db:"expert" json:"expert"`   
JobName       NullString       `db:"job_name" json:"job_name"`   
JobNumber       NullString       `db:"job_number" json:"job_number"`   
QueueId       NullInt64       `db:"queue_id" json:"queue_id"`   
}

type NQueueManager struct {
Id       int32       `db:"id" json:"id"`   
Name       string       `db:"name" json:"name"`   
Addr       string       `db:"addr" json:"addr"`   
CreateTime       string       `db:"create_time" json:"create_time"`   
LinkName       NullString       `db:"link_name" json:"link_name"`   
LinkHref       NullString       `db:"link_href" json:"link_href"`   
SmsAddr       NullString       `db:"sms_addr" json:"sms_addr"`   
CorpName       NullString       `db:"corp_name" json:"corp_name"`   
GiftPrice       NullFloat64       `db:"gift_price" json:"gift_price"`   
UpgradePrice       NullFloat64       `db:"upgrade_price" json:"upgrade_price"`   
GlassNum       NullInt64       `db:"glass_num" json:"glass_num"`   
TotalPrice       NullInt64       `db:"total_price" json:"total_price"`   
}

type NRedpackSignature struct {
Id       int32       `db:"id" json:"id"`   
SendName       NullString       `db:"send_name" json:"send_name"`   
TotalNum       NullString       `db:"total_num" json:"total_num"`   
Wishing       NullString       `db:"wishing" json:"wishing"`   
ActName       NullString       `db:"act_name" json:"act_name"`   
Remark       NullString       `db:"remark" json:"remark"`   
Phone       string       `db:"phone" json:"phone"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
Start       NullString       `db:"start" json:"start"`   
End       NullString       `db:"end" json:"end"`   
}

type NReturnOrder struct {
Id       int32       `db:"id" json:"id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ShipNo       NullString       `db:"ship_no" json:"ship_no"`   
LineOrderIds       NullString       `db:"line_order_ids" json:"line_order_ids"`   
UserPhone       NullString       `db:"user_phone" json:"user_phone"`   
Status       NullString       `db:"status" json:"status"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
}

type NReturnPrice struct {
Id       int32       `db:"id" json:"id"`   
Price       NullFloat64       `db:"price" json:"price"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
ClientId       NullInt64       `db:"client_id" json:"client_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Sn       NullString       `db:"sn" json:"sn"`   
Type       NullString       `db:"type" json:"type"`   
}

type NStockUser struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Title       NullString       `db:"title" json:"title"`   
CanIn       NullInt64       `db:"can_in" json:"can_in"`   
CanOut       NullInt64       `db:"can_out" json:"can_out"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Phone       NullString       `db:"phone" json:"phone"`   
Password       NullString       `db:"password" json:"password"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
BelongHouse       NullString       `db:"belong_house" json:"belong_house"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
IsRetire       NullInt64       `db:"is_retire" json:"is_retire"`   
IsAdmin       NullInt64       `db:"is_admin" json:"is_admin"`   
}

type NStoreHouseOrder struct {
Id       int32       `db:"id" json:"id"`   
FrameOrderId       int32       `db:"frame_order_id" json:"frame_order_id"`   
StoreHouse       NullString       `db:"store_house" json:"store_house"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
TotalPrice       NullFloat64       `db:"total_price" json:"total_price"`   
Status       NullString       `db:"status" json:"status"`   
StockUser       NullString       `db:"stock_user" json:"stock_user"`   
Third       string       `db:"third" json:"third"`   
Destination       NullString       `db:"destination" json:"destination"`   
Remark       NullString       `db:"remark" json:"remark"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Urgency       NullString       `db:"urgency" json:"urgency"`   
StockUserId       NullInt64       `db:"stock_user_id" json:"stock_user_id"`   
ThirdHandle       NullInt64       `db:"third_handle" json:"third_handle"`   
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
}

type NUser struct {
Id       int32       `db:"id" json:"id"`   
UserPhone       string       `db:"user_phone" json:"user_phone"`   
Sex       NullString       `db:"sex" json:"sex"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
Password       NullString       `db:"password" json:"password"`   
ShopId       NullString       `db:"shop_id" json:"shop_id"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
ReceiverPhone       NullString       `db:"receiver_phone" json:"receiver_phone"`   
ReceiverAddress       NullString       `db:"receiver_address" json:"receiver_address"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
LastLogin       NullTime       `db:"last_login" json:"last_login"`   // can be null, use with caution
VisitorSex       NullString       `db:"visitor_sex" json:"visitor_sex"`   
Corp       NullString       `db:"corp" json:"corp"`   
CorpEmail       NullString       `db:"corp_email" json:"corp_email"`   
RecommendCode       NullString       `db:"recommend_code" json:"recommend_code"`   
Role       NullString       `db:"role" json:"role"`   
Lens       NullString       `db:"lens" json:"lens"`   
Supplier       NullString       `db:"supplier" json:"supplier"`   
DealerStatus       NullString       `db:"dealer_status" json:"dealer_status"`   
IsSupplier       NullInt64       `db:"is_supplier" json:"is_supplier"`   
HeadImg       string       `db:"head_img" json:"head_img"`   
Name       NullString       `db:"name" json:"name"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Nickname       NullString       `db:"nickname" json:"nickname"`   
Age       NullString       `db:"age" json:"age"`   
}

type NWechatMini struct {
Id       int32       `db:"id" json:"id"`   
RefreshToken       NullString       `db:"refresh_token" json:"refresh_token"`   
Appid       NullString       `db:"appid" json:"appid"`   
NickName       NullString       `db:"nick_name" json:"nick_name"`   
HeadImg       NullString       `db:"head_img" json:"head_img"`   
ServiceTypeInfo       NullInt64       `db:"service_type_info" json:"service_type_info"`   
VerifyTypeInfo       NullInt64       `db:"verify_type_info" json:"verify_type_info"`   
OriginId       NullString       `db:"origin_id" json:"origin_id"`   
Alias       NullString       `db:"alias" json:"alias"`   
QrcodeUrl       NullString       `db:"qrcode_url" json:"qrcode_url"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
BusinessInfo       NullString       `db:"business_info" json:"business_info"`   
FuncInfo       NullString       `db:"func_info" json:"func_info"`   
Canceled       NullInt64       `db:"canceled" json:"canceled"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
SubMchId       NullString       `db:"sub_mch_id" json:"sub_mch_id"`   
PrincipalName       NullString       `db:"principal_name" json:"principal_name"`   
MManagement       NullInt64       `db:"m_management" json:"m_management"`   
MDevAnalyse       NullInt64       `db:"m_dev_analyse" json:"m_dev_analyse"`   
MKf       NullInt64       `db:"m_kf" json:"m_kf"`   
MOpen       NullInt64       `db:"m_open" json:"m_open"`   
MLogin       NullInt64       `db:"m_login" json:"m_login"`   
Version       NullString       `db:"version" json:"version"`   
OnlineTime       NullTime       `db:"online_time" json:"online_time"`   // can be null, use with caution
UserDesc       NullString       `db:"user_desc" json:"user_desc"`   
TemplateId       NullInt64       `db:"template_id" json:"template_id"`   
AuditId       NullInt64       `db:"audit_id" json:"audit_id"`   
AuditTime       NullTime       `db:"audit_time" json:"audit_time"`   // can be null, use with caution
}

type NWechatMiniHistory struct {
Id       int32       `db:"id" json:"id"`   
Appid       NullString       `db:"appid" json:"appid"`   
TemplateId       NullInt64       `db:"template_id" json:"template_id"`   
AuditId       NullInt64       `db:"audit_id" json:"audit_id"`   
UserDesc       NullString       `db:"user_desc" json:"user_desc"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
}

type NWechatOpen struct {
Id       int32       `db:"id" json:"id"`   
RefreshToken       NullString       `db:"refresh_token" json:"refresh_token"`   
Appid       NullString       `db:"appid" json:"appid"`   
NickName       NullString       `db:"nick_name" json:"nick_name"`   
HeadImg       NullString       `db:"head_img" json:"head_img"`   
ServiceTypeInfo       NullInt64       `db:"service_type_info" json:"service_type_info"`   
VerifyTypeInfo       NullInt64       `db:"verify_type_info" json:"verify_type_info"`   
OriginId       NullString       `db:"origin_id" json:"origin_id"`   
Alias       NullString       `db:"alias" json:"alias"`   
QrcodeUrl       NullString       `db:"qrcode_url" json:"qrcode_url"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
BusinessInfo       NullString       `db:"business_info" json:"business_info"`   
FuncInfo       NullString       `db:"func_info" json:"func_info"`   
Canceled       NullInt64       `db:"canceled" json:"canceled"`   
FMassSendNotice       NullInt64       `db:"f_mass_send_notice" json:"f_mass_send_notice"`   
FMsgManage       NullInt64       `db:"f_msg_manage" json:"f_msg_manage"`   
FMaterialManage       NullInt64       `db:"f_material_manage" json:"f_material_manage"`   
FUserManage       NullInt64       `db:"f_user_manage" json:"f_user_manage"`   
FAccountService       NullInt64       `db:"f_account_service" json:"f_account_service"`   
FWebpageService       NullInt64       `db:"f_webpage_service" json:"f_webpage_service"`   
FWechatTinyStore       NullInt64       `db:"f_wechat_tiny_store" json:"f_wechat_tiny_store"`   
FWechatMulCustomer       NullInt64       `db:"f_wechat_mul_customer" json:"f_wechat_mul_customer"`   
FWechatCardTicket       NullInt64       `db:"f_wechat_card_ticket" json:"f_wechat_card_ticket"`   
FWechatScan       NullInt64       `db:"f_wechat_scan" json:"f_wechat_scan"`   
FWechatWifi       NullInt64       `db:"f_wechat_wifi" json:"f_wechat_wifi"`   
FWechatShake       NullInt64       `db:"f_wechat_shake" json:"f_wechat_shake"`   
FWechatStore       NullInt64       `db:"f_wechat_store" json:"f_wechat_store"`   
FWechatPay       NullInt64       `db:"f_wechat_pay" json:"f_wechat_pay"`   
FCustomMenu       NullInt64       `db:"f_custom_menu" json:"f_custom_menu"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
SubMchId       NullString       `db:"sub_mch_id" json:"sub_mch_id"`   
PrincipalName       NullString       `db:"principal_name" json:"principal_name"`   
}

type NWechatOpenRefresh struct {
Id       int32       `db:"id" json:"id"`   
Appid       NullString       `db:"appid" json:"appid"`   
RefreshToken       NullString       `db:"refresh_token" json:"refresh_token"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NWechatOpenTemplate struct {
Id       int32       `db:"id" json:"id"`   
Appid       NullString       `db:"appid" json:"appid"`   
IdShort       NullString       `db:"id_short" json:"id_short"`   
TemplateId       NullString       `db:"template_id" json:"template_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type NWxApps struct {
Id       int32       `db:"id" json:"id"`   
App       NullString       `db:"app" json:"app"`   
AppId       NullString       `db:"app_id" json:"app_id"`   
AppSecret       NullString       `db:"app_secret" json:"app_secret"`   
PayId       NullString       `db:"pay_id" json:"pay_id"`   
PaySecret       NullString       `db:"pay_secret" json:"pay_secret"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type Order struct {
Id       int32       `db:"id" json:"id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
SphLeft       float64       `db:"sph_left" json:"sph_left"`   
SphRight       float64       `db:"sph_right" json:"sph_right"`   
CylLeft       NullFloat64       `db:"cyl_left" json:"cyl_left"`   
CylRight       NullFloat64       `db:"cyl_right" json:"cyl_right"`   
AxisLeft       NullInt64       `db:"axis_left" json:"axis_left"`   
AxisRight       NullInt64       `db:"axis_right" json:"axis_right"`   
Pd       int32       `db:"pd" json:"pd"`   
Remark       NullString       `db:"remark" json:"remark"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
}

type OrderPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
StatsEnabled       NullInt64       `db:"stats_enabled" json:"stats_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type PermitChange struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Operator       string       `db:"operator" json:"operator"`   
Event       string       `db:"event" json:"event"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type PermitChangeItems struct {
Id       int32       `db:"id" json:"id"`   
ChangeId       int32       `db:"change_id" json:"change_id"`   
System       string       `db:"system" json:"system"`   
Col       string       `db:"col" json:"col"`   
OldValue       string       `db:"old_value" json:"old_value"`   
NewValue       string       `db:"new_value" json:"new_value"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type RetailerGoods struct {
WxId       string       `db:"wx_id" json:"wx_id"`   
Phone       NullString       `db:"phone" json:"phone"`   
GoodsId       int32       `db:"goods_id" json:"goods_id"`   
Color       string       `db:"color" json:"color"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Id       int32       `db:"id" json:"id"`   
}

type ScoreConversionRecord struct {
Id       int32       `db:"id" json:"id"`   
BonusUserId       NullInt64       `db:"bonus_user_id" json:"bonus_user_id"`   
ShopGoodsId       NullInt64       `db:"shop_goods_id" json:"shop_goods_id"`   
GoodsName       NullString       `db:"goods_name" json:"goods_name"`   
Consignee       NullString       `db:"consignee" json:"consignee"`   
ConsigneePhone       NullString       `db:"consignee_phone" json:"consignee_phone"`   
Address       NullString       `db:"address" json:"address"`   
Score       int32       `db:"score" json:"score"`   
Status       NullInt64       `db:"status" json:"status"`   
IsMoney       NullInt64       `db:"is_money" json:"is_money"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
NeedCheck       NullInt64       `db:"need_check" json:"need_check"`   
BillNo       NullString       `db:"bill_no" json:"bill_no"`   
GoodsPrice       NullFloat64       `db:"goods_price" json:"goods_price"`   
CostScore       NullInt64       `db:"cost_score" json:"cost_score"`   
}

type Sequence struct {
Name       string       `db:"name" json:"name"`   
Value       int32       `db:"value" json:"value"`   
Increment       int32       `db:"increment" json:"increment"`   
}

type ShareFaceResult struct {
Id       int32       `db:"id" json:"id"`   
FileId       string       `db:"file_id" json:"file_id"`   
Width       int32       `db:"width" json:"width"`   
Height       int32       `db:"height" json:"height"`   
Result       string       `db:"result" json:"result"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
}

type ShopInfo struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
Address       NullString       `db:"address" json:"address"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
Area       NullString       `db:"area" json:"area"`   
Bank       NullString       `db:"bank" json:"bank"`   
BankNo       NullString       `db:"bank_no" json:"bank_no"`   
BankBranch       NullString       `db:"bank_branch" json:"bank_branch"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ActiveTags       NullString       `db:"active_tags" json:"active_tags"`   
MainImage       NullString       `db:"main_image" json:"main_image"`   
}

type ShopLens struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       string       `db:"name" json:"name"`   
OriginalPrice       float64       `db:"original_price" json:"original_price"`   
DiscountPrice       float64       `db:"discount_price" json:"discount_price"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDeleted       NullInt64       `db:"is_deleted" json:"is_deleted"`   
}

type ShopLikes struct {
Id       int32       `db:"id" json:"id"`   
UserId       NullInt64       `db:"user_id" json:"user_id"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
WxId       string       `db:"wx_id" json:"wx_id"`   
}

type ShopMember struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Name       string       `db:"name" json:"name"`   
Password       string       `db:"password" json:"password"`   
TelPhone       string       `db:"tel_phone" json:"tel_phone"`   
Store       NullString       `db:"store" json:"store"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type ShopOrder struct {
Id       int32       `db:"id" json:"id"`   
Sn       string       `db:"sn" json:"sn"`   
Phone       string       `db:"phone" json:"phone"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
UserName       NullString       `db:"user_name" json:"user_name"`   
Address       NullString       `db:"address" json:"address"`   
TotalNum       NullInt64       `db:"total_num" json:"total_num"`   
TotalPrice       NullFloat64       `db:"total_price" json:"total_price"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Status       NullString       `db:"status" json:"status"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
PayQuery       NullInt64       `db:"pay_query" json:"pay_query"`   
PayStatus       NullString       `db:"pay_status" json:"pay_status"`   
Remark       NullString       `db:"remark" json:"remark"`   
ShipNo       NullString       `db:"ship_no" json:"ship_no"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
Paid       NullInt64       `db:"paid" json:"paid"`   
ShipSelf       NullInt64       `db:"ship_self" json:"ship_self"`   
}

type ShopOrderLine struct {
Id       int32       `db:"id" json:"id"`   
ShopOrderId       NullInt64       `db:"shop_order_id" json:"shop_order_id"`   
GoodsId       NullInt64       `db:"goods_id" json:"goods_id"`   
GoodsName       NullString       `db:"goods_name" json:"goods_name"`   
ColorId       NullInt64       `db:"color_id" json:"color_id"`   
ColorName       NullString       `db:"color_name" json:"color_name"`   
Num       NullInt64       `db:"num" json:"num"`   
Price       NullFloat64       `db:"price" json:"price"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
Brand       NullString       `db:"brand" json:"brand"`   
Category       NullString       `db:"category" json:"category"`   
Sn       NullString       `db:"sn" json:"sn"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Collocations       NullString       `db:"collocations" json:"collocations"`   
}

type ShopPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
BuyEnabled       NullInt64       `db:"buy_enabled" json:"buy_enabled"`   
MemberEnabled       NullInt64       `db:"member_enabled" json:"member_enabled"`   
MainSubEnabled       NullInt64       `db:"main_sub_enabled" json:"main_sub_enabled"`   
SingleColumnEnabled       NullInt64       `db:"single_column_enabled" json:"single_column_enabled"`   
}

type ShopShowImages struct {
Id       int32       `db:"id" json:"id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
Img4       NullString       `db:"img_4" json:"img_4"`   
Img5       NullString       `db:"img_5" json:"img_5"`   
}

type ShopTag struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
ThirdId       int32       `db:"third_id" json:"third_id"`   
Img       NullString       `db:"img" json:"img"`   
IsShow       NullInt64       `db:"is_show" json:"is_show"`   
}

type Sid struct {
Id       int32       `db:"id" json:"id"`   
Name       string       `db:"name" json:"name"`   
Sid       int32       `db:"sid" json:"sid"`   
Phone       NullString       `db:"phone" json:"phone"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Progress       NullString       `db:"progress" json:"progress"`   
UThirdPhone       string       `db:"u_third_phone" json:"u_third_phone"`   
Hh       NullInt64       `db:"hh" json:"hh"`   
Xx       NullInt64       `db:"xx" json:"xx"`   
}

type SnImages struct {
Id       int32       `db:"id" json:"id"`   
Brand       string       `db:"brand" json:"brand"`   
Category       string       `db:"category" json:"category"`   
Sn       string       `db:"sn" json:"sn"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Img1       NullString       `db:"img_1" json:"img_1"`   
Img2       NullString       `db:"img_2" json:"img_2"`   
Img3       NullString       `db:"img_3" json:"img_3"`   
Img4       NullString       `db:"img_4" json:"img_4"`   
Img5       NullString       `db:"img_5" json:"img_5"`   
Img6       NullString       `db:"img_6" json:"img_6"`   
Img7       NullString       `db:"img_7" json:"img_7"`   
Img8       NullString       `db:"img_8" json:"img_8"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
}

type SocialLikes struct {
Id       int32       `db:"id" json:"id"`   
FileId       string       `db:"file_id" json:"file_id"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
Third       string       `db:"third" json:"third"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
FromUser       int32       `db:"from_user" json:"from_user"`   
ToUser       int32       `db:"to_user" json:"to_user"`   
OrderUsed       NullInt64       `db:"order_used" json:"order_used"`   
LikeType       int32       `db:"like_type" json:"like_type"`   
}

type SocialPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
WatermarkEnabled       NullInt64       `db:"watermark_enabled" json:"watermark_enabled"`   
ModelEnabled       NullInt64       `db:"model_enabled" json:"model_enabled"`   
StatsEnabled       NullInt64       `db:"stats_enabled" json:"stats_enabled"`   
SavelocalEnabled       NullInt64       `db:"savelocal_enabled" json:"savelocal_enabled"`   
TaobaoEnabled       NullInt64       `db:"taobao_enabled" json:"taobao_enabled"`   
HugebarginEnabled       NullInt64       `db:"hugebargin_enabled" json:"hugebargin_enabled"`   
BrandenterEnabled       NullInt64       `db:"brandenter_enabled" json:"brandenter_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
RankStep       NullInt64       `db:"rank_step" json:"rank_step"`   
LikesStep       NullInt64       `db:"likes_step" json:"likes_step"`   
RankEnabled       NullInt64       `db:"rank_enabled" json:"rank_enabled"`   
VoteEnabled       NullInt64       `db:"vote_enabled" json:"vote_enabled"`   
VoteStep       NullInt64       `db:"vote_step" json:"vote_step"`   
RateEnabled       NullInt64       `db:"rate_enabled" json:"rate_enabled"`   
WatermarkTime       NullString       `db:"watermark_time" json:"watermark_time"`   
ModelTime       NullString       `db:"model_time" json:"model_time"`   
StatsTime       NullString       `db:"stats_time" json:"stats_time"`   
SavelocalTime       NullString       `db:"savelocal_time" json:"savelocal_time"`   
TaobaoTime       NullString       `db:"taobao_time" json:"taobao_time"`   
HugebarginTime       NullString       `db:"hugebargin_time" json:"hugebargin_time"`   
BrandenterTime       NullString       `db:"brandenter_time" json:"brandenter_time"`   
RankTime       NullString       `db:"rank_time" json:"rank_time"`   
VoteTime       NullString       `db:"vote_time" json:"vote_time"`   
RateTime       NullString       `db:"rate_time" json:"rate_time"`   
BeginTime       NullString       `db:"begin_time" json:"begin_time"`   
}

type SocialPush struct {
Id       int32       `db:"id" json:"id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Likes       int32       `db:"likes" json:"likes"`   
Rank       int32       `db:"rank" json:"rank"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
RankUpdateTime       NullTime       `db:"rank_update_time" json:"rank_update_time"`   // can be null, use with caution
LikesUpdateTime       NullTime       `db:"likes_update_time" json:"likes_update_time"`   // can be null, use with caution
RankWaiting       int32       `db:"rank_waiting" json:"rank_waiting"`   
LikesWaiting       int32       `db:"likes_waiting" json:"likes_waiting"`   
}

type SocialUser struct {
Id       int32       `db:"id" json:"id"`   
WxId       string       `db:"wx_id" json:"wx_id"`   
Name       NullString       `db:"name" json:"name"`   
HeadImage       NullString       `db:"head_image" json:"head_image"`   
LastFileId       NullString       `db:"last_file_id" json:"last_file_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
BargainImg       NullString       `db:"bargain_img" json:"bargain_img"`   
TryImg       NullString       `db:"try_img" json:"try_img"`   
}

type SocialUserFace struct {
Id       int32       `db:"id" json:"id"`   
SocialUserId       int32       `db:"social_user_id" json:"social_user_id"`   
FaceId       string       `db:"face_id" json:"face_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Third       string       `db:"third" json:"third"`   
MatchImg       NullString       `db:"match_img" json:"match_img"`   
Src       NullString       `db:"src" json:"src"`   
}

type SocialVote struct {
Id       int32       `db:"id" json:"id"`   
SocialUserId       int32       `db:"social_user_id" json:"social_user_id"`   
Third       string       `db:"third" json:"third"`   
VoteImg       NullString       `db:"vote_img" json:"vote_img"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type SocialVoteLike struct {
Id       int32       `db:"id" json:"id"`   
TryCensusId       int32       `db:"try_census_id" json:"try_census_id"`   
FromUser       int32       `db:"from_user" json:"from_user"`   
ToUser       int32       `db:"to_user" json:"to_user"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type StockRights struct {
Id       int32       `db:"id" json:"id"`   
Departments       NullString       `db:"departments" json:"departments"`   
}

type StorePermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
OptometryDataEnabled       NullInt64       `db:"optometry_data_enabled" json:"optometry_data_enabled"`   
OrderRecordEnabled       NullInt64       `db:"order_record_enabled" json:"order_record_enabled"`   
ActiveNotificationEnabled       NullInt64       `db:"active_notification_enabled" json:"active_notification_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
SphEnabled       NullInt64       `db:"sph_enabled" json:"sph_enabled"`   
CylEnabled       NullInt64       `db:"cyl_enabled" json:"cyl_enabled"`   
AxisEnabled       NullInt64       `db:"axis_enabled" json:"axis_enabled"`   
PdEnabled       NullInt64       `db:"pd_enabled" json:"pd_enabled"`   
PhEnabled       NullInt64       `db:"ph_enabled" json:"ph_enabled"`   
CorctVisionEnabled       NullInt64       `db:"corct_vision_enabled" json:"corct_vision_enabled"`   
OriginVisionEnabled       NullInt64       `db:"origin_vision_enabled" json:"origin_vision_enabled"`   
PrismEnabled       NullInt64       `db:"prism_enabled" json:"prism_enabled"`   
StatusEnabled       NullInt64       `db:"status_enabled" json:"status_enabled"`   
AddEnabled       NullInt64       `db:"add_enabled" json:"add_enabled"`   
RemarkEnabled       NullInt64       `db:"remark_enabled" json:"remark_enabled"`   
TryOrderEnabled       NullInt64       `db:"try_order_enabled" json:"try_order_enabled"`   
EChannelEnabled       NullInt64       `db:"e_channel_enabled" json:"e_channel_enabled"`   
YxdEnabled       NullInt64       `db:"yxd_enabled" json:"yxd_enabled"`   
DomainEyeEnabled       NullInt64       `db:"domain_eye_enabled" json:"domain_eye_enabled"`   
UsedForEnabled       NullInt64       `db:"used_for_enabled" json:"used_for_enabled"`   
DataEnabled       NullInt64       `db:"data_enabled" json:"data_enabled"`   
DataSharedEnabled       NullInt64       `db:"data_shared_enabled" json:"data_shared_enabled"`   
SurveyEnabled       NullInt64       `db:"survey_enabled" json:"survey_enabled"`   
AbNewOptometryEnabled       NullInt64       `db:"ab_new_optometry_enabled" json:"ab_new_optometry_enabled"`   
QueueEnabled       NullInt64       `db:"queue_enabled" json:"queue_enabled"`   
GiftActivityEnabled       NullInt64       `db:"gift_activity_enabled" json:"gift_activity_enabled"`   
BargainEnabled       NullInt64       `db:"bargain_enabled" json:"bargain_enabled"`   
ScoreShopEnabled       NullInt64       `db:"score_shop_enabled" json:"score_shop_enabled"`   
ManufactureEnabled       NullInt64       `db:"manufacture_enabled" json:"manufacture_enabled"`   
BeginTime       NullString       `db:"begin_time" json:"begin_time"`   
OptometryDataTime       NullString       `db:"optometry_data_time" json:"optometry_data_time"`   
OrderRecordTime       NullString       `db:"order_record_time" json:"order_record_time"`   
ActiveNotificationTime       NullString       `db:"active_notification_time" json:"active_notification_time"`   
TryOrderTime       NullString       `db:"try_order_time" json:"try_order_time"`   
EChannelTime       NullString       `db:"e_channel_time" json:"e_channel_time"`   
YxdTime       NullString       `db:"yxd_time" json:"yxd_time"`   
DataTime       NullString       `db:"data_time" json:"data_time"`   
SurveyTime       NullString       `db:"survey_time" json:"survey_time"`   
GiftActivityTime       NullString       `db:"gift_activity_time" json:"gift_activity_time"`   
BargainTime       NullString       `db:"bargain_time" json:"bargain_time"`   
ScoreShopTime       NullString       `db:"score_shop_time" json:"score_shop_time"`   
ManufactureTime       NullString       `db:"manufacture_time" json:"manufacture_time"`   
OrderTypeEnabled       NullInt64       `db:"order_type_enabled" json:"order_type_enabled"`   
ChatManageEnabled       NullInt64       `db:"chat_manage_enabled" json:"chat_manage_enabled"`   
OrderTypeTime       NullString       `db:"order_type_time" json:"order_type_time"`   
ChatManageTime       NullString       `db:"chat_manage_time" json:"chat_manage_time"`   
CouponConsumeEnabled       NullInt64       `db:"coupon_consume_enabled" json:"coupon_consume_enabled"`   
CouponConsumeTime       NullString       `db:"coupon_consume_time" json:"coupon_consume_time"`   
}

type TPlatformWechatTemplate struct {
Id       int32       `db:"id" json:"id"`   
Appid       NullString       `db:"appid" json:"appid"`   
ShortId       NullString       `db:"short_id" json:"short_id"`   
TemplateId       NullString       `db:"template_id" json:"template_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TestBook struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
OriginPrice       NullFloat64       `db:"origin_price" json:"origin_price"`   
Price       NullFloat64       `db:"price" json:"price"`   
Stock       NullInt64       `db:"stock" json:"stock"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
AuthorId       NullInt64       `db:"author_id" json:"author_id"`   
}

type TestBookAuthor struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
Description       NullString       `db:"description" json:"description"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TestBookImage struct {
Id       int32       `db:"id" json:"id"`   
BookId       NullInt64       `db:"book_id" json:"book_id"`   
Pic       NullString       `db:"pic" json:"pic"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TryAttributes struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Sex       NullString       `db:"sex" json:"sex"`   
Frame       NullString       `db:"frame" json:"frame"`   
FrameStyle       NullString       `db:"frame_style" json:"frame_style"`   
Material       NullString       `db:"material" json:"material"`   
PriceRegion       NullString       `db:"price_region" json:"price_region"`   
FaceStyle       NullString       `db:"face_style" json:"face_style"`   
GlassType       NullString       `db:"glass_type" json:"glass_type"`   
}

type TryCensus struct {
Id       int32       `db:"id" json:"id"`   
Third       NullString       `db:"third" json:"third"`   
Type       NullInt64       `db:"type" json:"type"`   
StartTime       NullString       `db:"start_time" json:"start_time"`   
EndTime       NullString       `db:"end_time" json:"end_time"`   
Status       NullInt64       `db:"status" json:"status"`   
BeforeDay       NullInt64       `db:"before_day" json:"before_day"`   
AfterDay       NullInt64       `db:"after_day" json:"after_day"`   
ActivityInstruction       NullString       `db:"activity_instruction" json:"activity_instruction"`   
ShareTitle       NullString       `db:"share_title" json:"share_title"`   
UploadFinish       NullString       `db:"upload_finish" json:"upload_finish"`   
Entrance       NullString       `db:"entrance" json:"entrance"`   
Guide       NullString       `db:"guide" json:"guide"`   
Rules       NullString       `db:"rules" json:"rules"`   
}

type TryLoves struct {
Id       int32       `db:"id" json:"id"`   
UserId       string       `db:"user_id" json:"user_id"`   
ColorId       int32       `db:"color_id" json:"color_id"`   
Belongs       string       `db:"belongs" json:"belongs"`   
Type       NullString       `db:"type" json:"type"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TryModels struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
Type       NullString       `db:"type" json:"type"`   
Url       NullString       `db:"url" json:"url"`   
Position       NullString       `db:"position" json:"position"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
Width       int32       `db:"width" json:"width"`   
}

type TryPermit struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Granted       NullInt64       `db:"granted" json:"granted"`   
ExpireDate       NullTime       `db:"expire_date" json:"expire_date"`   
WatermarkEnabled       NullInt64       `db:"watermark_enabled" json:"watermark_enabled"`   
ModelEnabled       NullInt64       `db:"model_enabled" json:"model_enabled"`   
StatsEnabled       NullInt64       `db:"stats_enabled" json:"stats_enabled"`   
SavelocalEnabled       NullInt64       `db:"savelocal_enabled" json:"savelocal_enabled"`   
TaobaoEnabled       NullInt64       `db:"taobao_enabled" json:"taobao_enabled"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TryShareTemplate struct {
Id       int32       `db:"id" json:"id"`   
TemplateSn       NullString       `db:"template_sn" json:"template_sn"`   
TemplateName       NullString       `db:"template_name" json:"template_name"`   
Describe       NullString       `db:"describe" json:"describe"`   
ImgThumb       NullString       `db:"img_thumb" json:"img_thumb"`   
ImgDetail       NullString       `db:"img_detail" json:"img_detail"`   
ImgExplanationExit       NullInt64       `db:"img_explanation_exit" json:"img_explanation_exit"`   
Wishes       NullString       `db:"wishes" json:"wishes"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
BrandNameExit       NullInt64       `db:"brand_name_exit" json:"brand_name_exit"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
WishesExit       NullInt64       `db:"wishes_exit" json:"wishes_exit"`   
SceneSetting       NullString       `db:"scene_setting" json:"scene_setting"`   
}

type TryShareTemplateLine struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
TemplateId       int32       `db:"template_id" json:"template_id"`   
ImgWatermark       NullString       `db:"img_watermark" json:"img_watermark"`   
ImgExplanation       NullString       `db:"img_explanation" json:"img_explanation"`   
BrandName       NullString       `db:"brand_name" json:"brand_name"`   
Wishes       NullString       `db:"wishes" json:"wishes"`   
WishesSelected       NullInt64       `db:"wishes_selected" json:"wishes_selected"`   
SceneSetting       NullString       `db:"scene_setting" json:"scene_setting"`   
Enabled       NullInt64       `db:"enabled" json:"enabled"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type TryStore struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       NullString       `db:"name" json:"name"`   
TelPhone       NullString       `db:"tel_phone" json:"tel_phone"`   
StoreName       NullString       `db:"store_name" json:"store_name"`   
Address       NullString       `db:"address" json:"address"`   
Coordinate       NullString       `db:"coordinate" json:"coordinate"`   
IsThird       NullInt64       `db:"is_third" json:"is_third"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Deleted       int32       `db:"deleted" json:"deleted"`   
WaiterWxImg       NullString       `db:"waiter_wx_img" json:"waiter_wx_img"`   
MsgWxId       NullString       `db:"msg_wx_id" json:"msg_wx_id"`   
MsgHeadImg       NullString       `db:"msg_head_img" json:"msg_head_img"`   
MsgNickName       NullString       `db:"msg_nick_name" json:"msg_nick_name"`   
WaiterName       NullString       `db:"waiter_name" json:"waiter_name"`   
Province       NullString       `db:"province" json:"province"`   
City       NullString       `db:"city" json:"city"`   
District       NullString       `db:"district" json:"district"`   
Street       NullString       `db:"street" json:"street"`   
Type       NullString       `db:"type" json:"type"`   
StoreBalance       NullFloat64       `db:"store_balance" json:"store_balance"`   
CityCode       NullInt64       `db:"city_code" json:"city_code"`   
ShopNo       NullInt64       `db:"shop_no" json:"shop_no"`   
CreateAgentId       NullInt64       `db:"create_agent_id" json:"create_agent_id"`   
CreateStaffId       NullInt64       `db:"create_staff_id" json:"create_staff_id"`   
}

type TryStoreActivity struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       NullString       `db:"name" json:"name"`   
Detail       NullString       `db:"detail" json:"detail"`   
Img       NullString       `db:"img" json:"img"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Status       NullInt64       `db:"status" json:"status"`   
}

type UThird struct {
Id       int32       `db:"id" json:"id"`   
Phone       string       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Password       NullString       `db:"password" json:"password"`   
AssistPassword       NullString       `db:"assist_password" json:"assist_password"`   
Company       NullString       `db:"company" json:"company"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
UploadGranted       NullInt64       `db:"upload_granted" json:"upload_granted"`   
LensGranted       NullInt64       `db:"lens_granted" json:"lens_granted"`   
ImgLogo       NullString       `db:"img_logo" json:"img_logo"`   
BannerLogo       NullString       `db:"banner_logo" json:"banner_logo"`   
CompanyAddress       NullString       `db:"company_address" json:"company_address"`   
CompanyPhone       NullString       `db:"company_phone" json:"company_phone"`   
CompanyFax       NullString       `db:"company_fax" json:"company_fax"`   
CompanyBank       NullString       `db:"company_bank" json:"company_bank"`   
ImgWatermark       NullString       `db:"img_watermark" json:"img_watermark"`   
Model       NullString       `db:"model" json:"model"`   
Pos       NullString       `db:"pos" json:"pos"`   
HomeUrl       NullString       `db:"home_url" json:"home_url"`   
TryBtn       NullString       `db:"try_btn" json:"try_btn"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
BonusWx       NullString       `db:"bonus_wx" json:"bonus_wx"`   
BonusActivity       NullString       `db:"bonus_activity" json:"bonus_activity"`   
ScoreRmbRatio       int32       `db:"score_rmb_ratio" json:"score_rmb_ratio"`   
ConsumptionScoreRatio       int32       `db:"consumption_score_ratio" json:"consumption_score_ratio"`   
TryFavorScoreRatio       int32       `db:"try_favor_score_ratio" json:"try_favor_score_ratio"`   
MaxLikeNum       NullInt64       `db:"max_like_num" json:"max_like_num"`   
TryReduceMoney       NullFloat64       `db:"try_reduce_money" json:"try_reduce_money"`   
ConsumptionScoreEnabled       int32       `db:"consumption_score_enabled" json:"consumption_score_enabled"`   
TryFavorScoreEnabled       int32       `db:"try_favor_score_enabled" json:"try_favor_score_enabled"`   
TryRuleLink       NullString       `db:"try_rule_link" json:"try_rule_link"`   
Discount       NullFloat64       `db:"discount" json:"discount"`   
InitIncreaseBonus       NullInt64       `db:"init_increase_bonus" json:"init_increase_bonus"`   
MaxIncreaseBonus       NullInt64       `db:"max_increase_bonus" json:"max_increase_bonus"`   
MaxReduceBonus       NullInt64       `db:"max_reduce_bonus" json:"max_reduce_bonus"`   
LeastMoneyToUseBonus       NullInt64       `db:"least_money_to_use_bonus" json:"least_money_to_use_bonus"`   
DiscountEnabled       NullInt64       `db:"discount_enabled" json:"discount_enabled"`   
IncreaseBonus       NullInt64       `db:"increase_bonus" json:"increase_bonus"`   
ReduceMoney       NullInt64       `db:"reduce_money" json:"reduce_money"`   
ShareTemplate       NullInt64       `db:"share_template" json:"share_template"`   
Level       NullString       `db:"level" json:"level"`   
BonusEndUrl       NullString       `db:"bonus_end_url" json:"bonus_end_url"`   
Prerogative       NullString       `db:"prerogative" json:"prerogative"`   
Type       NullString       `db:"type" json:"type"`   
UseStatus       NullString       `db:"use_status" json:"use_status"`   
DealStatus       NullString       `db:"deal_status" json:"deal_status"`   
StoreNum       NullInt64       `db:"store_num" json:"store_num"`   
}

type UThirdClient struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Password       NullString       `db:"password" json:"password"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
BelongsName       NullString       `db:"belongs_name" json:"belongs_name"`   
Company       NullString       `db:"company" json:"company"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
ReceiverPhone       NullString       `db:"receiver_phone" json:"receiver_phone"`   
ReceiverAddr       NullString       `db:"receiver_addr" json:"receiver_addr"`   
ReceiverCity       NullString       `db:"receiver_city" json:"receiver_city"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
DisableCheck       NullInt64       `db:"disable_check" json:"disable_check"`   
IsDealer       NullInt64       `db:"is_dealer" json:"is_dealer"`   
BillPassword       NullString       `db:"bill_password" json:"bill_password"`   
IsReceive       NullInt64       `db:"is_receive" json:"is_receive"`   
}

type UThirdFrameClient struct {
Id       int32       `db:"id" json:"id"`   
Phone       NullString       `db:"phone" json:"phone"`   
Name       NullString       `db:"name" json:"name"`   
Password       NullString       `db:"password" json:"password"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
BelongsName       NullString       `db:"belongs_name" json:"belongs_name"`   
Store       NullString       `db:"store" json:"store"`   
ReceiverName       NullString       `db:"receiver_name" json:"receiver_name"`   
ReceiverPhone       NullString       `db:"receiver_phone" json:"receiver_phone"`   
ReceiverAddr       NullString       `db:"receiver_addr" json:"receiver_addr"`   
ReceiverCity       NullString       `db:"receiver_city" json:"receiver_city"`   
ShipCorp       NullString       `db:"ship_corp" json:"ship_corp"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
IsDealer       NullInt64       `db:"is_dealer" json:"is_dealer"`   
BillPassword       NullString       `db:"bill_password" json:"bill_password"`   
UThirdId       int32       `db:"u_third_id" json:"u_third_id"`   
PenddingBill       int32       `db:"pendding_bill" json:"pendding_bill"`   
}

type UThirdStaff struct {
Id       int32       `db:"id" json:"id"`   
Name       string       `db:"name" json:"name"`   
Phone       string       `db:"phone" json:"phone"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Password       string       `db:"password" json:"password"`   
Belongs       string       `db:"belongs" json:"belongs"`   
BelongsName       NullString       `db:"belongs_name" json:"belongs_name"`   
Store       NullString       `db:"store" json:"store"`   
IsRetire       int32       `db:"is_retire" json:"is_retire"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
}

type UThirdSupplier struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
SName       NullString       `db:"s_name" json:"s_name"`   
Address       NullString       `db:"address" json:"address"`   
Remark       NullString       `db:"remark" json:"remark"`   
Phone       NullString       `db:"phone" json:"phone"`   
}

type UThirdUser struct {
Id       int32       `db:"id" json:"id"`   
Name       NullString       `db:"name" json:"name"`   
WxId       NullString       `db:"wx_id" json:"wx_id"`   
Sid       NullInt64       `db:"sid" json:"sid"`   
Belongs       NullString       `db:"belongs" json:"belongs"`   
Phone       NullString       `db:"phone" json:"phone"`   
GlassName       NullString       `db:"glass_name" json:"glass_name"`   
DoorName       NullString       `db:"door_name" json:"door_name"`   
Price       NullString       `db:"price" json:"price"`   
Address       NullString       `db:"address" json:"address"`   
Discount       NullString       `db:"discount" json:"discount"`   
Remark       NullString       `db:"remark" json:"remark"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
QrUrl       NullString       `db:"qr_url" json:"qr_url"`   
ImgLogo       NullString       `db:"img_logo" json:"img_logo"`   
ImgWatermark       NullString       `db:"img_watermark" json:"img_watermark"`   
TelPhone       NullString       `db:"tel_phone" json:"tel_phone"`   
Coordinate       NullString       `db:"coordinate" json:"coordinate"`   
ActivityImg       NullString       `db:"activity_img" json:"activity_img"`   
Source       NullString       `db:"source" json:"source"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
IsThird       NullInt64       `db:"is_third" json:"is_third"`   
}

type UserAddress struct {
Id       int32       `db:"id" json:"id"`   
OptometryUserId       int32       `db:"optometry_user_id" json:"optometry_user_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
Name       NullString       `db:"name" json:"name"`   
Phone       NullString       `db:"phone" json:"phone"`   
Address       NullString       `db:"address" json:"address"`   
Deleted       NullInt64       `db:"deleted" json:"deleted"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
}

type VdataProvince struct {
Id       int32       `db:"id" json:"id"`   
Adcode       NullInt64       `db:"adcode" json:"adcode"`   
Lat       NullFloat64       `db:"lat" json:"lat"`   
Lng       NullFloat64       `db:"lng" json:"lng"`   
Province       NullString       `db:"province" json:"province"`   
Level       NullString       `db:"level" json:"level"`   
Parent       NullString       `db:"parent" json:"parent"`   
}

type VotePush struct {
Id       int32       `db:"id" json:"id"`   
UserId       int32       `db:"user_id" json:"user_id"`   
ThirdId       int32       `db:"third_id" json:"third_id"`   
VoteRank       int32       `db:"vote_rank" json:"vote_rank"`   
CreateTime       time.Time       `db:"create_time" json:"create_time"`   
UpdateTime       time.Time       `db:"update_time" json:"update_time"`   
RankUpdateTime       NullTime       `db:"rank_update_time" json:"rank_update_time"`   // can be null, use with caution
}

type WxInteractionLog struct {
Id       int32       `db:"id" json:"id"`   
ThirdId       NullInt64       `db:"third_id" json:"third_id"`   
Openid       NullString       `db:"openid" json:"openid"`   
EventKey       NullInt64       `db:"event_key" json:"event_key"`   
ChannelId       NullInt64       `db:"channel_id" json:"channel_id"`   
CreateTime       NullTime       `db:"create_time" json:"create_time"`   // can be null, use with caution
UpdateTime       NullTime       `db:"update_time" json:"update_time"`   // can be null, use with caution
Event       NullString       `db:"event" json:"event"`   
UserId       int32       `db:"user_id" json:"user_id"`   
StoreId       NullInt64       `db:"store_id" json:"store_id"`   
StaffId       NullInt64       `db:"staff_id" json:"staff_id"`   
Shared       NullInt64       `db:"shared" json:"shared"`   
}