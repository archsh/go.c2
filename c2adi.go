package c2

import (
	"encoding/xml"
	"fmt"
)

// the c2 adi definitions
/*
<?xml version="1.0" encoding="UTF-8"?>
<ADI xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<Objects>
		<Object ElementType="object_type" ID="object_id" Action="REGIST">
			<Property Name="property_name1">property_value1</Property>
			<Property Name="property_name2">property_value2</Property>
			<Property Name="property_name3">property_value3</Property>
		</Object>
    </Objects>
	<Mappings>
		<Mapping ID=”mapping_id” ParentType="parent_type" ParentID="parent_id" ElementType=”element_type” ElementID=”element_id” Action="REGIST">
			<Property name="property_name1">property_value1</Property>
			<Property name="property_name2">property_value2</Property>
		</Mapping>
	</Mappings>
</ADI>
*/

type Reply struct {
	Result      int    `xml:"Result"`
	Description string `xml:"Description"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Object struct {
	ID          string     `xml:",attr"`
	ElementType string     `xml:",attr"`
	Action      string     `xml:",attr"`
	Properties  []Property `xml:"Property"`
}

type Mapping struct {
	ID          string     `xml:",attr"`
	ParentType  string     `xml:",attr"`
	ParentID    string     `xml:",attr"`
	ElementType string     `xml:",attr"`
	ElementID   string     `xml:",attr"`
	Action      string     `xml:",attr"`
	Properties  []Property `xml:"Property"`
}

type ADI struct {
	XMLName  xml.Name  `xml:"http://www.w3.org/2001/XMLSchema-instance ADI"`
	Objects  []Object  `xml:"Objects>Object"`
	Mappings []Mapping `xml:"Mappings>Mapping"`
}

// Actions
const (
	REGIST = "REGIST"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// Element  Types
const (
	PROGRAM          = "Program"
	MOVIE            = "Movie"
	CAST             = "Cast"
	CAST_ROLE_MAP    = "CastRoleMap"
	CHANNEL          = "Channel"
	PHYSICAL_CHANNEL = "PhysicalChannel"
	SCHEDULE         = "Schedule"
	PICTURE          = "Picture"
	CATEGORY         = "Category"
	SERIES           = "Series"
	PACKAGE          = "Package"
)

const (
	STRING    = "String"
	DATE      = "Date"
	TIME      = "Time"
	DATETIME  = "Datetime"
	TIMESTAMP = "Timestamp"
	YEAR      = "Year"
	BOOLEAN   = "Boolean"
	INTEGER   = "Integer"
	FLOAT     = "Float"
)

type PropertyDescriptor struct {
	Name     string
	Type     string
	Desc     string
	Required bool
	Length   int
	Comment  string
	Default  interface{}
}

func Prop(n string, t string, d string, r bool, l int, v interface{}, c string) PropertyDescriptor {
	return PropertyDescriptor{
		Name: n, Type: t, Desc: d, Required: r, Length: l, Comment: c, Default: v,
	}
}

// Program定义
var ProgramProperties = []PropertyDescriptor{
	Prop("Name", STRING, "节目名称", true, 128, nil, ""),
	Prop("OrderNumber", STRING, "节目订购编号", false, 10, nil, "6位数字编号"),
	Prop("OriginalName", STRING, "原名", false, 128, nil, ""),
	Prop("SortName", STRING, "索引名称供界面排序", false, 128, nil, ""),
	Prop("SearchName", STRING, "搜索名称供界面搜索", false, 128, nil, ""),
	Prop("Genre", STRING, "Program的默认类别（Genre）", false, 128, nil, ""),
	Prop("ActorDisplay", STRING, "演员列表(只供显示)", false, 256, nil, ""),
	Prop("WriterDisplay", STRING, "作者列表(只供显示)", false, 256, nil, ""),
	Prop("OriginalCountry", STRING, "国家地区", false, 64, nil, ""),
	Prop("Language", STRING, "语言", false, 64, nil, ""),
	Prop("ReleaseYear", YEAR, "上映年份(YYYY)", false, 4, nil, ""),
	Prop("OrgAirDate", DATE, "首播日期(YYYYMMDD)", false, 8, nil, ""),
	Prop("LicensingWindowStart", TIMESTAMP, "有效开始时间(YYYYMMDDHH24MiSS)", false, 14, nil, ""),
	Prop("LicensingWindowEnd", TIMESTAMP, "有效结束时间(YYYYMMDDHH24MiSS)", false, 14, nil, ""),
	Prop("DisplayAsNew", INTEGER, "新到天数", false, 3, nil, ""),
	Prop("DisplayAsLastChance", INTEGER, "剩余天数", false, 3, nil, ""),
	Prop("Macrovision", BOOLEAN, "拷贝保护标志", false, 1, nil, " 0:无拷贝保护 1:有拷贝保护"),
	Prop("Description", STRING, "节目描述", false, 1024, nil, ""),
	Prop("PriceTaxIn", FLOAT, "列表定价", false, 14, nil, ""),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("SourceType", INTEGER, "1: VOD 5: Advertisement", false, 1, nil, ""),
	Prop("SeriesFlag", INTEGER, "0: 普通VOD 1: 连续剧剧集", false, 1, nil, ""),
	Prop("Type", STRING, "节目内容类型", true, 128, nil, "字符串，表示内容类型，例如：电影，连续剧，新闻，体育…"),
	Prop("Keywords", STRING, "关键字", false, 256, nil, "多个关键字之间使用分号分隔"),
	Prop("Tags", STRING, "关联标签", false, 256, nil, "多个标签之间使用分号分隔"),
	Prop("Reserve1", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve2", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve3", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve4", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve5", STRING, "保留字段", false, 256, nil, ""),
	Prop("StorageType", INTEGER, "存储分发策略要求", false, 2, nil, "0. 厂商CDN可不要存储本节目（在海量存储中保存，具体视频路径在Movie.OCSURL） 1. 厂商CDN存储本节目 >2. 自定义策略（具体对应策略在厂商系统中定义，可以做到部分节点覆盖，或者后拉视频文件…）;为海量存储，定制化存储增加，默认为1"),
	Prop("RMediaCode", STRING, "关联内容唯一标识", false, 128, nil, "为了支持表示不同屏的同一个内容关系"),
	Prop("Result", INTEGER, "错误代码", false, 1, 0, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
	Prop("DefinitionFlag", STRING, "节目清晰度标识", true, 1, nil, "标识内容节目是属于哪种清晰度（标清、高清、超清） 0：标清, 1：高清, 2：超清, 3. 4K, 4. 杜比"),
	Prop("DeviceGroupList", STRING, "设备类型组ID列表", false, 128, nil, "VOD、连续剧VOD支持的内容设备分组的ID列表，多个ID间使用英文逗号间"),
	Prop("BeginDuration", STRING, "片头时长", false, 8, nil, "（整形数值,单位::秒）"),
	Prop("EndDuration", STRING, "片尾时长", false, 8, nil, "（整形数值,单位::秒）"),
	Prop("PreDuration", STRING, "试看时长", false, 8, nil, "（整形数值,单位::秒）"),
}

// Movie定义
var MovieProperties = []PropertyDescriptor{
	Prop("Type", INTEGER, "媒体类型", true, 1, nil, "1:正片 2:预览片"),
	Prop("FileURL", STRING, "媒体文件URL", true, 1024, nil, "ftp://username:password@ip:port/...  标准FTP协议"),
	Prop("SourceDRMType", INTEGER, "0: No DRM 1: BES DRM", true, 1, nil, ""),
	Prop("DestDRMType", INTEGER, "0: No DRM 1: BES DRM", true, 1, nil, ""),
	Prop("AudioType", INTEGER, "0: 其他, 1: Monaural 单声道, 2: Stereo 多声道, 3: Two-nation monaural 双单声道, 4: Two-nation stereo 双多声道, 5: AC3(5:1 channel) AC3声道", true, 2, nil, ""),
	Prop("ScreenFormat", INTEGER, "0: 4x3, 1: 16x9(Wide)", true, 1, nil, ""),
	Prop("ClosedCaptioning", INTEGER, "字幕标志", true, 1, nil, "0:无字幕, 1:有字幕"),
	Prop("OCSURL", STRING, "在海量存储中的视频URL", true, 1024, nil, "类似 Rtsp://ip:port/1/2/3.ts"),
	Prop("Duration", STRING, "播放时长HHMISSFF", true, 12, nil, "（时分秒帧）"),
	Prop("FileSize", INTEGER, "文件大小，单位为Byte", true, 16, nil, ""),
	Prop("BitRateType", INTEGER, "码流", true, 3, nil, "0:其他\n1:400k\n2:700k\n3:1.3M\n4:2M\n5:2.5M\n6:8M\n7:10M\n8:12M\n9:16M\n51：1.3M（标清VBR）\n52：2M（标清VBR）\n53：4M（高清VBR）\n54：6M（超清VBR）\n55：2.3M（高清VBR）\n56：25M（4K H264 VBR）\n57：14M（4K H264 VBR）\n58：15M（4K H265 VBR）\n59：8M（4K H265 VBR）\n\n310: 18M（4K VBR）\n311: 21M（4K VBR）"),
	Prop("VideoType", INTEGER, "编码格式", true, 3, nil, "1:H.264\n2:MPEG4\n3:AVS\n4:MPEG2\n5:MP3\n7:H.265"),
	Prop("AudioFormat", INTEGER, "编码格式", true, 3, nil, "1. MP2\n    2. AAC\n    3. AMR"),
	Prop("Resolution", INTEGER, "分辨率类型", true, 3, nil, "1:QCIF\n2:QVGA\n3:2/3 D1\n4:3/4 D1\n5:D1\n6:720P\n7:1080i\n8:1080P\n9:4K(3840*2160)\n10:4K(4096*2160)"),
	Prop("VideoProfile", INTEGER, "Video Profile", true, 3, nil, "1:QCIF\n2:QVGA\n3:2/3 D1\n4:3/4 D1\n5:D1\n6:720P\n7:1080i\n8:1080P\n9:4K(3840*2160)\n10:4K(4096*2160)"),
	Prop("SystemLayer", INTEGER, "System Layer", true, 3, nil, "1:TS\n2:3GP\n3:MP4"),
	Prop("ServiceType", INTEGER, "服务类型", true, 3, nil, "0x01：在线播放(默认)\n0x10：支持下载\n0X11:  在线播放+下载"),
	Prop("MediaSepc", STRING, "媒资视音频编转码信息", true, 128, nil, "建议按照央视规范，对Movie添加视音频编转码信息; CMS直接透传此字段给EPG管理门户"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Cast定义
var CastProperties = []PropertyDescriptor{
	Prop("Name", STRING, "人物名称", false, 64, nil, ""),
	Prop("PersonDisplayName", STRING, "显示名称", false, 64, nil, ""),
	Prop("PersonSortName", STRING, "排序名称", false, 64, nil, ""),
	Prop("PersonSearchName", STRING, "索引名称", false, 64, nil, ""),
	Prop("FirstName", STRING, "姓", true, 32, nil, ""),
	Prop("MiddleName", STRING, "中间名", false, 32, nil, ""),
	Prop("LastName", STRING, "名", false, 32, nil, ""),
	Prop("Sex", INTEGER, "性别", false, 1, nil, "0:女 1:男"),
	Prop("Birthday", DATE, "生日", false, 16, nil, ""),
	Prop("Hometown", STRING, "籍贯", false, 128, nil, ""),
	Prop("Education", STRING, "教育程度", false, 128, nil, ""),
	Prop("Height", INTEGER, "身高", false, 5, nil, ""),
	Prop("Weight", INTEGER, "体重", false, 5, nil, ""),
	Prop("BloodGroup", STRING, "血型", false, 2, nil, ""),
	Prop("Marriage", BOOLEAN, "婚", false, 1, nil, "0: 未婚\n1: 已婚"),
	Prop("Favorite", STRING, "爱好", false, 128, nil, ""),
	Prop("Webpage", STRING, "主页", false, 128, nil, ""),
	Prop("Description", STRING, "描述信息", false, 1024, nil, ""),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// CastRoleMap定义
var CastRoleMapProperties = []PropertyDescriptor{
	Prop("CastRole", STRING, "演职角色名称", true, 32, nil, ""),
	Prop("CastID", STRING, "人物ID", true, 32, nil, ""),
	Prop("CastCode", STRING, "人物Code", true, 128, nil, "表示关联的Cast的标识"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Channel定义
var ChannelProperties = []PropertyDescriptor{
	Prop("ChannelNumber", INTEGER, "建议频道号", false, 3, nil, "可以不填， \n如果发送方送的ChannelNumber为空，接收方填充自己默认值，如果发送方送的ChannelNumber跟接受方有冲突，接收方也填充自己的默认值，如果不冲突则按照发送方的值填充！"),
	Prop("Name", STRING, "频道名称", true, 64, nil, ""),
	Prop("CallSign", STRING, "台标名称", true, 10, nil, ""),
	Prop("TimeShift", BOOLEAN, "时移标志", false, 1, nil, "0:不生效 1:生效"),
	Prop("StorageDuration", INTEGER, "存储时长", false, 9, nil, "单位小时,仅仅对Timeshift有效"),
	Prop("TimeShiftDuration", INTEGER, "默认时移时长", false, 9, nil, "单位分钟\n(Reserved),仅仅对Timeshift有效"),
	Prop("Description", STRING, "描述信息", false, 1024, nil, ""),
	Prop("Country", STRING, "国家", false, 10, nil, ""),
	Prop("State", STRING, "州/省", false, 10, nil, ""),
	Prop("City", STRING, "城市", false, 10, nil, ""),
	Prop("ZipCode", STRING, "邮编", false, 10, nil, ""),
	Prop("Type", INTEGER, "频道类型", true, 1, nil, "1:直播频道"),
	Prop("SubType", INTEGER, "频道 类型", false, 1, nil, "当Type为1(直播频道)\n1: 信号源来自live\n2: 信号源来自virtual"),
	Prop("Language", STRING, "语言", false, 128, nil, ""),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("StartTime", TIME, "播放开始时间", true, 4, nil, "(HH24MI)"),
	Prop("EndTime", TIME, "播放结束时间", true, 4, nil, "(HH24MI)"),
	Prop("Macrovision", BOOLEAN, "拷贝保护标志", false, 1, nil, " 0:无拷贝保护 1:有拷贝保护"),
	Prop("VideoType", STRING, "视频参数", false, 10, nil, ""),
	Prop("AudioType", STRING, "音频参数", false, 10, nil, ""),
	Prop("StreamType", STRING, "码流标志", false, 1, nil, ""),
	Prop("Bilingual", STRING, "双语标志", false, 1, nil, ""),
	Prop("URL", STRING, "Web频道入口地址", false, 128, nil, "当type=2时，这个属性必填。"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// PhysicalChannel定义
var PhysicalChannelProperties = []PropertyDescriptor{
	Prop("ChannelID", STRING, "频道ID", true, 32, nil, ""),
	Prop("ChannelCode", STRING, "频道Code", true, 128, nil, "表示关联的Channel的标识"),
	Prop("BitRateType", INTEGER, "码流", true, 1, nil, "2: 2M\n4: 4M"),
	Prop("MultiCastIP", STRING, "组播IP", true, 64, nil, ""),
	Prop("MultiCastPort", INTEGER, "组播端口", true, 5, nil, ""),
	Prop("BitrateCount", INTEGER, "码率个数", false, 2, nil, "当BizDomain为PC时必填，范围1～10; 支持发布WEBTV的直播新增属性"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Schedule定义
var ScheduleProperties = []PropertyDescriptor{
	Prop("ChannelID", STRING, "频道ID", true, 32, nil, ""),
	Prop("ChannelCode", STRING, "频道Code", true, 128, nil, "表示关联的Channel的标识"),
	Prop("ProgramName", STRING, "节目名称", true, 128, nil, ""),
	Prop("SearchName", STRING, "搜索名称供界面搜索", false, 128, nil, "对应到Program(TVOD)的SearchName"),
	Prop("Genre", STRING, "Schedule的默认类别（Genre）", false, 128, nil, "对应到Program(TVOD)的Genre"),
	Prop("SourceType", INTEGER, "节目订购编号", false, 1, nil, "1: VOD\n5: Advertisement; 对应到Program(TVOD)的SourceType"),
	Prop("StartDate", DATE, "节目开播日期", true, 8, nil, "(YYYYMMDD)"),
	Prop("StartTime", TIME, "节目开播时间", true, 6, nil, "(HH24MISS)"),
	Prop("Duration", TIME, "节目时长", true, 6, nil, "(HH24MISS)"),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("Description", STRING, "描述信息", false, 1024, nil, ""),
	Prop("ObjectType", INTEGER, "关联的对象类型", false, 1, nil, "1：LiveTV Program(直播频道用)\t\n2：VOD Program(虚拟频道用)\t\n3：LiveTV Channel(虚拟频道中引入的直播频道)"),
	Prop("ObjectCode", STRING, "关联的对象Code", false, 128, nil, "ObjectType为1时，填ProgramCode(对于LIVE流，原来没有相关的Program关联，需新增Program)\nObjectType为2时，填ProgramCode(关联已有VOD)\nObjectType为3时，填ChannelCode(关联已有LiveChannel)"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Picture定义
var PictureProperties = []PropertyDescriptor{
	Prop("FileURL", STRING, "图片文件URL", false, 1024, nil, ""),
	Prop("Description", STRING, "描述信息", false, 1024, nil, ""),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Category定义
var CategoryProperties = []PropertyDescriptor{
	Prop("Name", STRING, "分类名称", true, 64, nil, ""),
	Prop("Sequence", INTEGER, "显示顺序号", true, 3, nil, ""),
	Prop("ParentID", STRING, "父节点ID", true, 32, nil, "若是根节点，ParentID为0"),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("Description", STRING, "描述信息", false, 1024, nil, ""),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Series定义
var SeriesProperties = []PropertyDescriptor{
	Prop("Name", STRING, "连续剧名称", true, 128, nil, ""),
	Prop("OrderNumber", INTEGER, "订购编号", false, 10, nil, ""),
	Prop("OriginalName", STRING, "原名", false, 128, nil, ""),
	Prop("SortName", STRING, "排序名称", false, 128, nil, ""),
	Prop("SearchName", STRING, "索引名称", false, 160, nil, ""),
	Prop("OrgAirDate", DATE, "首播日期(YYYYMMDD)", false, 8, nil, ""),
	Prop("LicensingWindowStart", TIMESTAMP, "有效定购开始时间(YYYYMMDDHH24MiSS)", true, 14, nil, ""),
	Prop("LicensingWindowEnd", TIMESTAMP, "有效定购结束时间(YYYYMMDDHH24MiSS)", true, 14, nil, ""),
	Prop("DisplayAsNew", INTEGER, "新到天数", false, 3, nil, ""),
	Prop("DisplayAsLastChance", INTEGER, "剩余天数", false, 3, nil, ""),
	Prop("Macrovision", BOOLEAN, "拷贝保护标志", false, 1, nil, " 0:无拷贝保护 1:有拷贝保护"),
	Prop("Price", FLOAT, "含税定价", false, 14, nil, ""),
	Prop("VolumnCount", INTEGER, "总集数", false, 5, nil, ""),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("Description", STRING, "节目描述", false, 1024, nil, ""),
	Prop("Type", STRING, "节目内容类型", true, 128, nil, "字符串，表示内容类型，例如：电影，连续剧，新闻，体育…"),
	Prop("Keywords", STRING, "关键字", false, 256, nil, "多个关键字之间使用分号分隔"),
	Prop("Tags", STRING, "关联标签", false, 256, nil, "多个标签之间使用分号分隔"),
	Prop("Reserve1", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve2", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve3", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve4", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve5", STRING, "保留字段", false, 256, nil, ""),
	Prop("RMediaCode", STRING, "关联内容唯一标识", false, 128, nil, "为了支持表示不同屏的同一个内容关系"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
	Prop("DeviceGroupList", STRING, "设备类型组ID列表", false, 128, nil, "VOD、连续剧VOD支持的内容设备分组的ID列表，多个ID间使用英文逗号间"),
}

// Package定义
var PackageProperties = []PropertyDescriptor{
	Prop("Name", STRING, "Package名称", true, 64, nil, ""),
	Prop("Type", INTEGER, "Package类型", true, 2, nil, "0:   VOD包\n2:   Channel包\n3：TVOD\n4:   SVOD\n5：PVOD包 \n\n99: Mix(待定义)"),
	Prop("SortName", STRING, "排序名称", false, 128, nil, ""),
	Prop("SearchName", STRING, "索引名称", false, 160, nil, ""),
	Prop("RentalPeriod", INTEGER, "租用有效期", false, 9, nil, "(小时)"),
	Prop("OrderNumber", STRING, "定购编号", true, 32, nil, ""),
	Prop("LicensingWindowStart", TIMESTAMP, "有效定购开始时间(YYYYMMDDHH24MiSS)", true, 14, nil, ""),
	Prop("LicensingWindowEnd", TIMESTAMP, "有效定购结束时间(YYYYMMDDHH24MiSS)", true, 14, nil, ""),
	Prop("Price", FLOAT, "含税定价", false, 14, nil, ""),
	Prop("Status", BOOLEAN, "状态标志", true, 1, nil, " 0:失效 1:生效"),
	Prop("Description", STRING, "节目描述", false, 1024, nil, ""),
	Prop("Keywords", STRING, "关键字", false, 256, nil, "多个关键字之间使用分号分隔"),
	Prop("Tags", STRING, "关联标签", false, 256, nil, "多个标签之间使用分号分隔"),
	Prop("Reserve1", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve2", STRING, "保留字段", false, 1024, nil, ""),
	Prop("Reserve3", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve4", STRING, "保留字段", false, 256, nil, ""),
	Prop("Reserve5", STRING, "保留字段", false, 256, nil, ""),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

// Mappings定义
var MappringProperties = []PropertyDescriptor{
	Prop("Type", INTEGER, "映射时的类型", false, 2, nil, "当Mapping的ParentType为Picture时：\n0: 缩略图\n1: 海报\n2: 剧照\n3: 图标\n4: 标题图\n5: 广告图\n6: 草图\n7: 背景图\n9: 频道图片\n10: 频道黑白图片\n11: 频道Logo\n12: 频道名字图片\n99: 其他"),
	Prop("Sequence", INTEGER, "序列号", false, 10, nil, "当Mapping关系涉及Picture时，此字段为必填，展示顺序有上有平台保证；\n当Mapping关系涉及Series和Program间绑定时，Sequence必须填写；\n\n说明：\n1、绑定海报时长度最大支持3位；\n2、主子集绑定时长度最大支持五位；  \n3、内容绑定栏目时长度最大支持7位"),
	Prop("ValidStart", TIMESTAMP, "服务起始时间(YYYYMMDDHH24MiSS)", true, 14, nil, "当Mapping的ParentType为SVOD时, 标识SVOD节目的服务起始时间\n (YYYYMMDDHH24MiSS)"),
	Prop("ValidEnd", TIMESTAMP, "服务终止时间(YYYYMMDDHH24MiSS)", true, 14, nil, "当Mapping的ParentType为SVOD时, 标识SVOD节目的服务终止时间\n(YYYYMMDDHH24MiSS)"),
	Prop("Result", INTEGER, "错误代码", false, 1, nil, "应答文件包含,0: 成功 其他: 错误代码"),
	Prop("ErrorDescription", STRING, "错误描述", false, 1024, nil, "应答文件包含"),
}

var ObjectProperties map[string][]PropertyDescriptor = map[string][]PropertyDescriptor{
	PROGRAM:          ProgramProperties,
	MOVIE:            MovieProperties,
	SERIES:           SeriesProperties,
	CATEGORY:         CategoryProperties,
	CAST:             CastProperties,
	CAST_ROLE_MAP:    CastRoleMapProperties,
	CHANNEL:          ChannelProperties,
	PHYSICAL_CHANNEL: PhysicalChannelProperties,
	PACKAGE:          PackageProperties,
	SCHEDULE:         ScheduleProperties,
	PICTURE:          PictureProperties,
}

func l2m(pds []PropertyDescriptor) map[string]PropertyDescriptor {
	var m = make(map[string]PropertyDescriptor, len(pds))
	for _, p := range pds {
		m[p.Name] = p
	}

	return m
}

var ObjectPropertyMaps map[string]map[string]PropertyDescriptor = map[string]map[string]PropertyDescriptor{
	PROGRAM:          l2m(ProgramProperties),
	MOVIE:            l2m(MovieProperties),
	SERIES:           l2m(SeriesProperties),
	CATEGORY:         l2m(CategoryProperties),
	CAST:             l2m(CastProperties),
	CAST_ROLE_MAP:    l2m(CastRoleMapProperties),
	CHANNEL:          l2m(ChannelProperties),
	PHYSICAL_CHANNEL: l2m(PhysicalChannelProperties),
	PACKAGE:          l2m(PackageProperties),
	SCHEDULE:         l2m(ScheduleProperties),
	PICTURE:          l2m(PictureProperties),
}

func NewObject(tp string, id string, action string, properties ...Property) Object {
	var pdmap map[string]PropertyDescriptor
	if pds, b := ObjectPropertyMaps[tp]; !b {
		panic("invalid type:" + tp)
	} else {
		pdmap = pds
	}
	var o = Object{
		ElementType: tp,
		ID:          id,
		Action:      action,
	}
	var pm = make(map[string]Property, len(properties))
	for _, p := range properties {
		pm[p.Name] = p
	}
	for k, p := range pdmap {
		if pp, b := pm[k]; b {
			o.Properties = append(o.Properties, Property{Name: k, Value: pp.Value})
		} else if nil != p.Default {
			o.Properties = append(o.Properties, Property{Name: k, Value: fmt.Sprint(p.Default)})
		}
	}
	return o
}
