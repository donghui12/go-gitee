/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 获取目录Tree
type Tree struct {
	Sha string `json:"sha,omitempty"`
	Url string `json:"url,omitempty"`
	Tree string `json:"tree,omitempty"`
	Truncated string `json:"truncated,omitempty"`
}
