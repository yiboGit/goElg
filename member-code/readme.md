* n_wechat_open --- 查出app_id，third_id 公众号授权
* 批量查出公众的会员卡 card_id  接口 使用app_id
* 依据 card_id 查出  会员卡的MemberCard ---id 存e_member_card
* 获取用户已领取卡券接口 找到memberCode
* e_optometry_user  openid == wx_id 
`phone_wx & wx_id & !member_code
for {
    limit 50
    break
}
`

