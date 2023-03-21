package DeviceCore

type StructLayout struct {
	Id           int    `json:"id"`            // 设备ID
	Name         string `json:"name"`          // 设备名称
	Ip           string `json:"ip"`            // 控制IP
	Port         string `json:"port"`          // 端口
	Prev         int    `json:"prev"`          // 前一个设备ID
	Next         []int  `json:"next"`          // 下一个设备ID
	DeliveryTime int64  `json:"delivery_time"` // 接收上一个设备传递所需时间
}

func InitLayout() (layout []StructLayout) {
	return layout
}

// 保存设备布局
func (layout *StructLayout) Save() (err error) {
	return err
}

// 获取设备布局
func (layout *StructLayout) GetLayout() (err error) {
	return err
}

// 获取设备布局
func (layout *StructLayout) AddLayout() (err error) {
	return err
}
