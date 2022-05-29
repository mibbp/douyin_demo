package controller

var DemoVideos = []Video{
	{
		Id:     1,
		Author: DemoUser,
		//PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		PlayUrl: "http://vfx.mtime.cn/Video/2019/02/04/mp4/190204084208765161.mp4",

		//CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CoverUrl:      "https://s2.loli.net/2022/05/22/hB36N84btKpcPQT.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:   1,
		User: DemoUser,
		//User: TestUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
		//Content:    "feed test",
		//CreateDate: "05-22",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "zhanglei",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
var TestUser = User{
	Id: 2,
	//Name:          "TestUser",
	Name:          "mibbp",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      true,
}
