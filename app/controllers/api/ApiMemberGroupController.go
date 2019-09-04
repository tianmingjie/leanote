package api

import (
	"github.com/leanote/leanote/app/info"
	"github.com/revel/revel"
)

// 标签API

type ApiMemberGroup struct {
	ApiBaseContrller
}

func (c ApiMemberGroup) CreateGroup(title string) revel.Result {

	admin := userService.GetUserInfoByAny("admin")
	re := info.NewRe()
	re.Ok, re.Item = groupService.AddGroup(admin.UserId.Hex(), title)
	return c.RenderJSON(re)
}

// 获取同步的标签
// [OK]
// > afterUsn的笔记
// 返回 {ChunkHighUsn: 本下最大的usn, 借此可以知道是否还有, Notebooks: []}
func (c ApiMemberGroup) GetGroup(title string) revel.Result {

	ret := groupService.GetGroupByTitle(title);
	return c.RenderJSON(ret)
}


func (c ApiMemberGroup) GetGroups() revel.Result {

	ret := groupService.GetAllGroups();
	return c.RenderJSON(ret)
}

// 添加Tag
// [OK]
// 不会产生冲突, 即使里面有
// 返回
/*
{
  "TagId": "551978dd99c37b9bc5000001",
  "UserId": "54a1676399c37b1c77000002",
  "Tag": "32",
  "Usn": 25,
  "Count": 1,
  "CreatedTime": "2015-03-31T00:25:01.149312407+08:00",
  "UpdatedTime": "2015-03-31T00:25:01.149312407+08:00",
  "IsDeleted": false
}
*/
func (c ApiMemberGroup) AddUser(title string,email string) revel.Result {

	admin := userService.GetUserInfoByAny("admin")
	userInfo := userService.GetUserInfoByAny(email)
	groupInfo:=groupService.GetGroupByTitle(title)
	ret, _ := groupService.AddUser(admin.UserId.Hex(), groupInfo.GroupId.Hex(), userInfo.UserId.Hex())

	return c.RenderJSON(ret)
}

// 删除标签
// [OK]
func (c ApiMemberGroup) DeleteUser(title string,email string) revel.Result {
	admin := userService.GetUserInfoByAny("admin")
	userInfo := userService.GetUserInfoByAny(email)
	groupInfo:=groupService.GetGroupByTitle(title)
	ret, _ := groupService.DeleteUser(admin.UserId.Hex(), groupInfo.GroupId.Hex(), userInfo.UserId.Hex())
	return c.RenderJSON(ret)
}
