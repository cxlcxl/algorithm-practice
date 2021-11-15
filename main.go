package main

func main() {
	//c := colly.NewCollector(
	//	colly.Debugger(&debug.LogDebugger{}),
	//	//colly.MaxDepth(1),
	//)
	//c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4503.5 Safari/537.36"
	//// 所有a标签，上设置回调函数
	//c.OnRequest(func(r *colly.Request) {
	//	r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	//	r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//	r.Headers.Set("Accept", "application/json, text/plain, */*")
	//	r.Headers.Set("Connection", "keep-alive")
	//	r.Headers.Set("Content-Type", "application/json")
	//	//r.Headers.Set("Referer", "https://ads.huawei.com/")
	//	r.Headers.Set("Host", "svc-dra.ads.huawei.com")
	//	r.Headers.Set("frontEndLang", "zh_CN")
	//	r.Headers.Set("Sec-Fetch-Mode", "cors")
	//	r.Headers.Set("Sec-Fetch-Site", "same-site")
	//	r.Headers.Set("Sec-Fetch-Dest", "empty")
	//	r.Headers.Set("sec-ch-ua-mobile", "?0")
	//	r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	//	r.Headers.Set("Cookie", "corpId=494043463919075840; userId=494043463919075840; ppsdspportal=B84979AFA8838BF7C21C8E976165DACB36A4ED34521B4EE9; channel_name=google; channel_category=search; _ga=GA1.2.919820308.1630980754; _gcl_au=1.1.1789065317.1630980754; _dmpa_id=394230bad4173badd4dc6103823391631072837181.1631072768.1.1631072768.1631072768; _dmpa_ref=[\"\",\"\",1631072837,\"https://consumer.huawei.com/cn/phones/p50-pro/specs/\"]; _fbp=fb.1.1631072903905.1807093554; Hm_lvt_48e5a2ca327922f1ee2bb5ea69bdd0a6=1630981020,1631072570,1631072644,1632797317; utag_main=v_id:017bbe0748350019cc5600a5621703071001a06900bd0$_sn:2$_se:4$_ss:0$_st:1632799123582$dc_visit:2$ses_id:1632797316749;exp-session$_pn:1;exp-session$dc_event:3;exp-session$dc_region:ap-northeast-1;exp-session; Hm_lpvt_48e5a2ca327922f1ee2bb5ea69bdd0a6=1632797380; language=en_US; HuaweiID_CAS_ISCASLOGIN=true; CASLOGINSITE=1; LOGINACCSITE=5; dsp_ab_flag=new; authInfo=\"{\\\"expiretime\\\":\\\"20211109T064953Z\\\",\\\"rtCiphertext\\\":\\\"rcH6zCa3bLiYzQs6AEIqNI89QDqDtSyCT3ZA6+KTgR9qpVnBFCo42fS51UuRywOAngU9GGpa6NRQkAeQEVOkZDXrDb8tOCFpJQFWawvTaAM44/kPXu5ZivoZjqtPRkAbGXbrWrmk63xiMv3MbNJKdQUuOqSzr8xIpz6sqER9oto6QaxaVNCJdurAcGQ5TuRA\\\",\\\"createtime\\\":\\\"20211109T054953Z\\\",\\\"signature\\\":\\\"d0f9655581245aa78a7588ba21ec9667f87324b63bb6eb263d192c3816b92991\\\",\\\"accesstoken\\\":\\\"CwEAAAAAdcJzCtPWyXiIUpXwRLUrZJeQ7Ef8nR6FxwomLN4fyqiXifWr0v/qa6WH0M2DM1CtzAxutgTtVgOWa5SHbYB/Tjhuc2vk+jRaRwYlu40gKJS6jxY=\\\",\\\"siteID\\\":\\\"5\\\"}\"; csrfToken=93B61682444F3E27C1FC770206A4F3F36D1F03F20B21D5DC1C; USERMGT_SID=e8a440a5334328ee9099fd0b8196a6f132c728b2c245d8e42e8b3ce1b397e4988a81926871642b812c83517a9e33b6808f0d200efdf75c7aec7a9cc59133f757; dsp_siteId=3")
	//})
	//c.OnResponse(func(response *colly.Response) {
	//	fmt.Println("response：", string(response.Body))
	//})
	//c.OnError(func(resp *colly.Response, errHttp error) {
	//	fmt.Println(errHttp)
	//})
	////params := map[string]interface{}{
	////	"flowResource":           1,
	////	"isQueryDeletedAudience": false,
	////	"promotionType":          2,
	////	"targetingList":          []string{"languageList", "dmpGeo", "currentDmpGeo", "levelGeo", "currentLevelGeo", "dmpGender", "dmpAge", "appCategory", "appInterest", "modelType", "devicePrice", "networkType", "installedApps", "notInstalledApps", "preAudience", "notPreAudience", "mac", "activeDays", "pushModelType", "locationCategory"},
	////}
	////bytes, _ := json.Marshal(params)
	//err := c.PostRaw("https://svc-dra.ads.huawei.com/ppsdspportal/v1/csrftoken/query", nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(bytes))
	//err = c.Visit("https://svc-dra.ads.huawei.com/ppsdspportal/v1/promotion_type/query")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// http post 请求
	//params := map[string]interface{}{
	//	"flowResource":           1,
	//	"isQueryDeletedAudience": false,
	//	"promotionType":          2,
	//	"targetingList":          []string{"languageList", "dmpGeo", "currentDmpGeo", "levelGeo", "currentLevelGeo", "dmpGender", "dmpAge", "appCategory", "appInterest", "modelType", "devicePrice", "networkType", "installedApps", "notInstalledApps", "preAudience", "notPreAudience", "mac", "activeDays", "pushModelType", "locationCategory"},
	//}
	//paramsBytes, err := json.Marshal(params)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println(string(paramsBytes))
	//req, err := http.NewRequest("POST", "https://svc-dra.ads.huawei.com/ppsdspportal/v1/targeting/query?showLoading=false", bytes.NewReader(paramsBytes))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//curl := http.Client{}
	//res, err := curl.Do(req)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(body, res.Body)
}
