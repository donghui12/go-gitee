/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 获取一个企业
type EnterpriseBasic struct {
	Id        int32  `json:"id,omitempty"`
	Path      string `json:"path,omitempty"`
	Name      string `json:"name,omitempty"`
	Url       string `json:"url,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}
