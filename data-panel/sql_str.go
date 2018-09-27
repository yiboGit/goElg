package main

// GetCustomerData 使用的sql (客户数)
var (
	CustomerCaseWhen = `case when province in('广东省','海南省','广西壮族自治区','香港特别行政区','澳门特别行政区','台湾省') then '华南地区' 
	when province in('北京市','天津市','河北省','山西省','内蒙古自治区') then '华北地区' 
	when province in('上海市','浙江省','安徽省','江苏省','山东省','福建省','江西省') then '华东地区' 
	when province in('河南省','湖北省','湖南省') then '华中地区' 
	when province in('四川省','云南省','贵州省','西藏自治区','重庆市') then '西南地区' 
	when province in('陕西省','甘肃省','青海省','宁夏回族自治区','新疆维吾尔自治区') then '西北地区' 
	when province in('黑龙江省','吉林省','辽宁省') then '东北地区' 
	else '未填写' end`
	CustomerSql = `select %s as 'Area',
	count(distinct id) as 'CusCreateNum',
	sum(store_num) as 'StoreCreateNum'
	from u_third
	where
	create_time between '%s' and '%s'
	and id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	group by %s
	order by 'Area'
	asc limit 10000
	`
	CustomerTypeSql = `select type, count(distinct id) as "CusCreateNum"
	from u_third
	where
	create_time between '%s' and '%s'
	and id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	group by type
	order by type
	asc limit 10000
	`
)

// 门店数
var (
	StoreCaseWhen = `case when province in('广东省','海南省','广西壮族自治区','香港特别行政区','澳门特别行政区','台湾省') then '华南地区' 
	when province in('北京市','天津市','河北省','山西省','内蒙古自治区') then '华北地区' 
	when province in('上海市','浙江省','安徽省','江苏省','山东省','福建省','江西省') then '华东地区' 
	when province in('河南省','湖北省','湖南省') then '华中地区' 
	when province in('四川省','云南省','贵州省','西藏自治区','重庆市') then '西南地区' 
	when province in('陕西省','甘肃省','青海省','宁夏回族自治区','新疆维吾尔自治区') then '西北地区' 
	when province in('黑龙江省','吉林省','辽宁省') then '东北地区' 
	else '未填写' end
	`
	StoreSql = `select %s as 'Area',
	count(distinct id) as 'StoreNum'
	from try_store
	where
	create_time between '%s' and '%s'
	and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and deleted<>'1'
	group by 
	%s
	`

	StoreCenterSql = `select if(c.center_name is null,'无加工中心',c.center_name),count(distinct s.id)  
	from try_store s left join m_center_store cs  on s.id = cs.store_id left join m_center c on c.id = cs.center_id 
	where 
	s.create_time between '%s' and '%s'
	and s.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and s.deleted<>'1'
	group by  if(c.center_name is null,'无加工中心',c.center_name)`
)

var (
	EmployeeSql = `select count(distinct id) as 'EmployeeNum'
	from e_staff
	where 
	create_time between '%s' and '%s'
	and deleted<>1
	and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	`
)

// 会员数
var (
	MemberSql = `select count(distinct case when is_subscribe = '1' then id end) as 'BeenFans',
	count(distinct case when is_subscribe=0 and unsubscribe_time is not null then id end) as 'UnFans',
	count(distinct case when card_status='activation' then id end) as 'MembershipCard',
	count(distinct case when card_status='activation' and shop_mini_openid is not null then id end) as 'Mall',
	count(distinct case when phone is not null and phone_wx is null then id end) as "Records",
	count(*) as 'TotalMemberNum'
	from e_optometry_user
	where
	create_time between '%s' and '%s'
	and deleted<>1
	and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	`
)

// 验光记录
var (
	OptometrySql = `
	select (count(case when (sph_left<=0 and sph_left >-3.5) then id end) + count(case when (sph_right<=0 and sph_right >-3.5) then id end)) as 'OneLevel',
		(count(case when (sph_left <=-3.5 and sph_left >-5.5)  then id end) + count(case when (sph_right <=-3.5 and sph_right >-5.5)  then id end)) as 'TwoLevel',
		(count(case when (sph_left <=-5.5)  then id end) + count(case when (sph_right <=-5.5)  then id end)) as 'ThreeLevel',
		(count(case when sph_left >0 then id end)+count(case when sph_right >0 then id end)) as 'PlusLevel',
		(count(*)*2) as 'OptometryNum'
		from e_optometry
		where 
		create_time between '%s' and '%s'
		and deleted<>1
		and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
		'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	`
)

// GetExpensesRecord 使用sql (消费记录/消费额)
var (
	ConsumptionSql = `select goods_title as 'PayMethod',
	count(distinct id) as 'RecordNum',
		sum(total_fee) as 'RealPay',
		sum(total_fee) as 'TotalPrice'
		from e_pay_order
		where status='success'
	and create_time between '%s' and '%s'
		and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
		'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	group by goods_title`
	ExpensesCaseWhen = `case when pay_method='cash' then '现金支付'
	when pay_method='card' then '银行卡支付'
	when pay_method='wxPay' then '微信记账'
	when pay_method='aliPay' then '支付宝记账'
	when pay_method='recharge' then '充值卡支付'
	when pay_method='eglassWxPay' then '微信在线支付'
	when pay_method='eglassAliPay' then '支付宝在线支付'
	when pay_method='xiaoruiPay' then '小瑞免息分期'
	else '未分类' end`
	ExpensesSql = `select %s as 'PayMethod',
count(distinct id) as 'RecordNum',
sum(price-discount_price-card_price-try_reduce_money) as 'RealPay',
sum(price) as 'TotalPrice'
from e_consumption
where 
create_time between '%s' and '%s'
and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
and paid='1'
and (deleted<>'1' or disabled<>'1')
group by
%s	
`
)

var (
	DepositCaseWhen = `case when b.pay_channel='cash' then '现金支付' 
	when b.pay_channel='card' then '银行卡支付'
	when b.pay_channel='wxPay' then '微信记账'
	when b.pay_channel='aliPay' then '支付宝记账'
	when b.pay_channel='eglassWxPay' then '微信在线支付'
	when b.pay_channel='eglassAliPay' then '支付宝在线支付'
	when b.pay_channel='storeCardPay' then '门店台卡支付'
	else 'null' end`
	DepositSql = `
	select %s as 'PayMethod',
count(distinct a.id) as 'PayRecordNum',
sum(b.money) as 'PayMoney'
from
(
select
*
from
e_recharge_record
where
third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
and deleted<>'1'
and create_time between '%s' and '%s'
)a join
(
select
*
from
e_recharge_pay
where
third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
and result_code='success'
and create_time between '%s' and '%s'
)b on a.id=b.recharge_record_id
group by
%s`
)

var (
	AllCardSql = `select count(distinct id) as 'TotalCardNum'
	from
	wx_card_log
	where
	third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and card_type='coupon'
	and event in('send_card','user_get_card','user_gifting_card','user_consume_card')
	and create_time between '%s' and '%s'
	`

	CardSql = `select %s as 'CardEvent',
	count(distinct id) as 'CardNum'
	from
	wx_card_log
	where
	third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224','225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and card_type='coupon'
	and event in('send_card','user_get_card','user_gifting_card','user_consume_card')
	and create_time between '%s' and '%s'
	group by
	%s
	`
	CardCaseWhen = `case when event='send_card' then '发送卡券'
	when event='user_get_card' then '领取卡券'
	when event='user_gifting_card' then '转赠卡券'
	when event='user_consume_card' then '核销卡券'
	else 'null' end
	`
)

var (
	// 全部的活动
	ActivityNumSql = `select  if(substr(end_time,1,10)<=curdate(),'已结束','进行中') as 'ActivityState',
	count(id) as 'ActivityNum'
	   from e_activity
	   where
	   type='fans'
	   and third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	   '225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	   and create_time between '%s' and '%s'
   group by 
   if(substr(end_time,1,10)<=curdate(),'已结束','进行中')`

	ChannelNumSql = `select  if(substr(a.end_time,1,10)<=curdate(),'已结束','进行中') as 'ActivityState',
	count(b.id) as 'EChannelNum'
	   from e_activity a left join e_channel b on a.id=b.activity_id
	   where
	   a.type='fans'
	   and a.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	   '225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	   and a.create_time between '%s' and '%s'
   group by 
   if(substr(a.end_time,1,10)<=curdate(),'已结束','进行中')`

	FansNumSql = `select if(substr(b.end_time,1,10)<=curdate(),'已结束','进行中') as 'ActivityState',
   count(distinct case when substr(c.create_time,1,16)=substr(b.create_time,1,16) then c.id end) as 'ActFansNum'
   from (
   select a.end_time as end_time,b.id as id,b.create_time as create_time,b.user_id as user_id
	   from e_activity a left join e_channel b on a.id=b.activity_id
	   where
	   a.type='fans'
	   and a.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	   '225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	   and a.create_time between '%s' and '%s') b left join  e_optometry_user c 
   on b.user_id=c.id
   group by 
   if(substr(b.end_time,1,10)<=curdate(),'已结束','进行中')
   `
)

var (
	PosterNumSql = `select count(a.id) as 'ProsterCresteNum'
	from e_poster a
	left join e_poster_release b 
	on a.id=b.poster_id
	where 
	a.deleted<>'1'
	and a.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and a.create_time between '%s' and '%s'
	`

	ChanNumSql = `select count(c.id) as 'ChanCreateNum'
	from e_poster a
	left join e_channel c
	on a.id=c.poster_id
	where 
	a.deleted<>'1'
	and a.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and a.create_time between '%s' and '%s'`

	FanNumSql = `select count(d.id) as 'NewFansNum' from e_optometry_user d right join
	(
	select c.id 
	from e_poster a
	left join e_channel c
	on a.id=c.poster_id
	where 
	a.deleted<>'1'
	and a.third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and a.create_time between '%s' and '%s') e
	on e.id=d.channel_id`
)

// 商品数量
var (
	GoodsLibSql = `select count(id) as 'GoodsLibNum' from e_goods_lib 
	where deleted <>1
	and create_time between '%s' and '%s'`
	ScoreGoodsSql = `select count(id) as 'ScoreGoodNum' from e_score_goods 
	where third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and deleted<>'1'
	and create_time between '%s' and '%s'`
	ShopCartSql = `select sum(total_num) as 'ShopCartNum' from e_shop_cart
	where third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and deleted<>'1'
	and create_time between '%s' and '%s'`
)

//
var (
	PayStatusSql = `select case when pay_status='paying' then '待付款订单'
	when pay_status='paid' then '已付款订单'
	when pay_status='canceled' then '已取消订单'
	else null end as 'OrderStatus',
	count(id) as 'OrderNum',
	sum(money_amount) as 'OrderMoney'
	from e_score_order
	where 
	third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and deleted<>'1'
	and create_time between '%s' and '%s'
	group by 
	case when pay_status='paying' then '待付款订单'
	when pay_status='paid' then '已付款订单'
	when pay_status='canceled' then '已取消订单'
	else null end
	`

	CheckStatusSql = `select case when check_status='uncheck' then '未核销订单'
	when check_status='checked' then '已核销订单'
	else null end as 'OrderStatus',
	count(id) as 'OrderNum',
	sum(money_amount) as 'OrderMoney'
	from e_score_order
	where 
	third_id not in('5','16','17','92','95','106','113','124','128','136','196','201','209','213','224',
	'225','233','234','235','256','271','324','325','326','363','365','367','376','414','417','570')
	and deleted<>'1'
	and create_time between '%s' and '%s'
	group by 
	case when check_status='uncheck' then '未核销订单'
	when check_status='checked' then '已核销订单'
	else null end`
)

var (
	SourceOrderSql = `select case when source in('PJDS') then '线下加工单'
	when source in('JXE','JXE_DIRECT') then '线上加工单'
	else 'null' end as 'Classification',
	count(id) as 'WorkOrderNum'
	from m_process 
	where create_time between '%s' and '%s'
	and status not in('9','10','11','cancelPay')
	and deleted<>'1'
	group by case when source in('PJDS') then '线下加工单'
	when source in('JXE','JXE_DIRECT') then '线上加工单'
	else 'null' end
	`

	CenterOrderSql = `select b.center_name as 'Classification',
	count(a.id) as 'WorkOrderNum'
	from m_process a left join m_center b on a.center_id=b.id
	where a.create_time between '%s' and '%s'
	and a.status not in('9','10','11','cancelPay')
	and a.deleted<>'1'
	group by b.center_name
	order by count(a.id) DESC
	limit 10
	`

	StatusOrderSql = `select case when status='0' then '待付款'
	when status='1' then '取料中'
	when status='2' then '待加工'
	when status='3' then '加工中'
	when status='4' then '质检中'
	when status='5' then '待送货'
	when status='6' then '送货中'
	when status='7' then '待确认'
	when status='8' then '已完成'
	when status='9' then '已取消'
	when status='10' then '已超时'
	when status='11' then '备料中'
	when status='cancelPay' then '已取消支付'
	else null end as 'Classification',
	count(id) as 'WorkOrderNum'
	from m_process 
	where create_time between '%s' and '%s'
	and deleted<>'1'
	group by case when status='0' then '待付款'
	when status='1' then '取料中'
	when status='2' then '待加工'
	when status='3' then '加工中'
	when status='4' then '质检中'
	when status='5' then '待送货'
	when status='6' then '送货中'
	when status='7' then '待确认'
	when status='8' then '已完成'
	when status='9' then '已取消'
	when status='10' then '已超时'
	when status='11' then '备料中'
	when status='cancelPay' then '已取消支付'
	else null end`
)

var (
	CenterDeliverSql = `select b.center_name as 'Classification',
	 count(a.id) as 'DeliverOrderNum'
	 from m_deliver a left join m_center b on a.center_id=b.id
	 where a.deleted<>'1'
	 and a.create_time between '%s' and '%s'
	 group by b.center_name
	 order by count(a.id) DESC
	 limit 10`

	StatusDeliverSql = `select case when order_payment_status='0' then '未支付' 
	 when order_payment_status='1' then '已支付'
	 when order_payment_status='2' then '已取消' 
	 else null end as 'Classification',
	 count(id) as 'DeliverOrderNum'
	 from m_deliver
	 where deleted<>'1'
	 and create_time between '%s' and '%s'
	 group by case when order_payment_status='0' then '未支付' 
	 when order_payment_status='1' then '已支付'
	 when order_payment_status='2' then '已取消' 
	 else null end
	 `
)

var (
	BrandSalesSql = `select b.brand as 'Classification',
	(count(a.left_info)+count(a.right_info)) as 'GoodsNum'
	from m_process a join m_process_goods b on a.id=b.process_id 
	where 
	b.category = 'len'
	and a.deleted<>'1'
	and a.status in('3','4','5', '6', '7', '8')
	and a.create_time between '%s' and '%s'
	group by b.brand
	order by (count(a.left_info)+count(a.right_info)) DESC
	limit 10`

	CenterSalesSql = `select c.center_name as 'Classification',
	(count(a.left_info)+count(a.right_info)) as 'GoodsNum'
	from m_process a join m_process_goods b on a.id=b.process_id left join m_center c on a.center_id=c.id
	where 
	b.category = 'len'
	and a.deleted<>'1'
	and a.status in('3','4','5', '6', '7', '8')
	and a.create_time between '%s' and '%s'
	group by c.center_name
	order by (count(a.left_info)+count(a.right_info)) DESC
	`
)
