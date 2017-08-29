package models

var OneNews News

func init() {
	OneNews = News{"10086", "1503982613", "又一支神秘部队长官公开亮相", "https://p3.pstatp.com/large/37dc0003fd4001769fcf", Desc{"10086", "火箭军洲际战略导弹旅“东风第一旅”旅长王锡民刚刚讲完，解放军海军辽宁舰舰长刘喆就登台开讲，为大家讲述了辽宁舰的故事"}}
}

type News struct {
	Id        string
	Timestamp string
	Title     string
	Imag      string
	Desc      Desc
}

type Desc struct {
	Id      string
	Content string
}
