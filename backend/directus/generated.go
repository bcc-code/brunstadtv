// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package directus

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// CreateAssetCreate_assets_itemAssets includes the requested fields of the GraphQL type assets.
type CreateAssetCreate_assets_itemAssets struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

// GetId returns CreateAssetCreate_assets_itemAssets.Id, and is useful for accessing the field via an interface.
func (v *CreateAssetCreate_assets_itemAssets) GetId() string { return v.Id }

// GetName returns CreateAssetCreate_assets_itemAssets.Name, and is useful for accessing the field via an interface.
func (v *CreateAssetCreate_assets_itemAssets) GetName() string { return v.Name }

// GetDuration returns CreateAssetCreate_assets_itemAssets.Duration, and is useful for accessing the field via an interface.
func (v *CreateAssetCreate_assets_itemAssets) GetDuration() string { return v.Duration }

// CreateAssetResponse is returned by CreateAsset on success.
type CreateAssetResponse struct {
	Create_assets_item *CreateAssetCreate_assets_itemAssets `json:"create_assets_item"`
}

// GetCreate_assets_item returns CreateAssetResponse.Create_assets_item, and is useful for accessing the field via an interface.
func (v *CreateAssetResponse) GetCreate_assets_item() *CreateAssetCreate_assets_itemAssets {
	return v.Create_assets_item
}

type Create_assetfiles_input struct {
	Id                string                       `json:"id"`
	User_created      *Create_directus_users_input `json:"user_created,omitempty"`
	Date_created      time.Time                    `json:"date_created"`
	Date_created_func *Datetime_functionsInput     `json:"date_created_func,omitempty"`
	User_updated      *Create_directus_users_input `json:"user_updated,omitempty"`
	Date_updated      time.Time                    `json:"date_updated"`
	Date_updated_func *Datetime_functionsInput     `json:"date_updated_func,omitempty"`
	Filename          string                       `json:"filename"`
	Type              string                       `json:"type"`
	Asset_id          *Create_assets_input         `json:"asset_id,omitempty"`
	Service           string                       `json:"service"`
	Metadata          string                       `json:"metadata"`
}

// GetId returns Create_assetfiles_input.Id, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetId() string { return v.Id }

// GetUser_created returns Create_assetfiles_input.User_created, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetUser_created() *Create_directus_users_input {
	return v.User_created
}

// GetDate_created returns Create_assetfiles_input.Date_created, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetDate_created() time.Time { return v.Date_created }

// GetDate_created_func returns Create_assetfiles_input.Date_created_func, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetDate_created_func() *Datetime_functionsInput {
	return v.Date_created_func
}

// GetUser_updated returns Create_assetfiles_input.User_updated, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetUser_updated() *Create_directus_users_input {
	return v.User_updated
}

// GetDate_updated returns Create_assetfiles_input.Date_updated, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetDate_updated() time.Time { return v.Date_updated }

// GetDate_updated_func returns Create_assetfiles_input.Date_updated_func, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetDate_updated_func() *Datetime_functionsInput {
	return v.Date_updated_func
}

// GetFilename returns Create_assetfiles_input.Filename, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetFilename() string { return v.Filename }

// GetType returns Create_assetfiles_input.Type, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetType() string { return v.Type }

// GetAsset_id returns Create_assetfiles_input.Asset_id, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetAsset_id() *Create_assets_input { return v.Asset_id }

// GetService returns Create_assetfiles_input.Service, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetService() string { return v.Service }

// GetMetadata returns Create_assetfiles_input.Metadata, and is useful for accessing the field via an interface.
func (v *Create_assetfiles_input) GetMetadata() string { return v.Metadata }

type Create_assets_input struct {
	Id                string                       `json:"id"`
	User_created      *Create_directus_users_input `json:"user_created,omitempty"`
	Date_created      time.Time                    `json:"date_created"`
	Date_created_func *Datetime_functionsInput     `json:"date_created_func,omitempty"`
	User_updated      *Create_directus_users_input `json:"user_updated,omitempty"`
	Date_updated      time.Time                    `json:"date_updated"`
	Date_updated_func *Datetime_functionsInput     `json:"date_updated_func,omitempty"`
	Name              string                       `json:"name"`
	Duration          string                       `json:"duration"`
	Files             []*Create_assetfiles_input   `json:"files,omitempty"`
}

// GetId returns Create_assets_input.Id, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetId() string { return v.Id }

// GetUser_created returns Create_assets_input.User_created, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetUser_created() *Create_directus_users_input { return v.User_created }

// GetDate_created returns Create_assets_input.Date_created, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetDate_created() time.Time { return v.Date_created }

// GetDate_created_func returns Create_assets_input.Date_created_func, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetDate_created_func() *Datetime_functionsInput {
	return v.Date_created_func
}

// GetUser_updated returns Create_assets_input.User_updated, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetUser_updated() *Create_directus_users_input { return v.User_updated }

// GetDate_updated returns Create_assets_input.Date_updated, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetDate_updated() time.Time { return v.Date_updated }

// GetDate_updated_func returns Create_assets_input.Date_updated_func, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetDate_updated_func() *Datetime_functionsInput {
	return v.Date_updated_func
}

// GetName returns Create_assets_input.Name, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetName() string { return v.Name }

// GetDuration returns Create_assets_input.Duration, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetDuration() string { return v.Duration }

// GetFiles returns Create_assets_input.Files, and is useful for accessing the field via an interface.
func (v *Create_assets_input) GetFiles() []*Create_assetfiles_input { return v.Files }

type Create_directus_files_input struct {
	Id                string                         `json:"id"`
	Storage           string                         `json:"storage"`
	Filename_disk     string                         `json:"filename_disk"`
	Filename_download string                         `json:"filename_download"`
	Title             string                         `json:"title"`
	Type              string                         `json:"type"`
	Folder            *Create_directus_folders_input `json:"folder,omitempty"`
	Uploaded_by       *Create_directus_users_input   `json:"uploaded_by,omitempty"`
	Uploaded_on       time.Time                      `json:"uploaded_on"`
	Uploaded_on_func  *Datetime_functionsInput       `json:"uploaded_on_func,omitempty"`
	Modified_by       *Create_directus_users_input   `json:"modified_by,omitempty"`
	Modified_on       time.Time                      `json:"modified_on"`
	Modified_on_func  *Datetime_functionsInput       `json:"modified_on_func,omitempty"`
	Charset           string                         `json:"charset"`
	Filesize          int                            `json:"filesize"`
	Width             int                            `json:"width"`
	Height            int                            `json:"height"`
	Duration          int                            `json:"duration"`
	Embed             string                         `json:"embed"`
	Description       string                         `json:"description"`
	Location          string                         `json:"location"`
	Tags              string                         `json:"tags"`
	Metadata          string                         `json:"metadata"`
}

// GetId returns Create_directus_files_input.Id, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetId() string { return v.Id }

// GetStorage returns Create_directus_files_input.Storage, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetStorage() string { return v.Storage }

// GetFilename_disk returns Create_directus_files_input.Filename_disk, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetFilename_disk() string { return v.Filename_disk }

// GetFilename_download returns Create_directus_files_input.Filename_download, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetFilename_download() string { return v.Filename_download }

// GetTitle returns Create_directus_files_input.Title, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetTitle() string { return v.Title }

// GetType returns Create_directus_files_input.Type, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetType() string { return v.Type }

// GetFolder returns Create_directus_files_input.Folder, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetFolder() *Create_directus_folders_input { return v.Folder }

// GetUploaded_by returns Create_directus_files_input.Uploaded_by, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetUploaded_by() *Create_directus_users_input {
	return v.Uploaded_by
}

// GetUploaded_on returns Create_directus_files_input.Uploaded_on, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetUploaded_on() time.Time { return v.Uploaded_on }

// GetUploaded_on_func returns Create_directus_files_input.Uploaded_on_func, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetUploaded_on_func() *Datetime_functionsInput {
	return v.Uploaded_on_func
}

// GetModified_by returns Create_directus_files_input.Modified_by, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetModified_by() *Create_directus_users_input {
	return v.Modified_by
}

// GetModified_on returns Create_directus_files_input.Modified_on, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetModified_on() time.Time { return v.Modified_on }

// GetModified_on_func returns Create_directus_files_input.Modified_on_func, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetModified_on_func() *Datetime_functionsInput {
	return v.Modified_on_func
}

// GetCharset returns Create_directus_files_input.Charset, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetCharset() string { return v.Charset }

// GetFilesize returns Create_directus_files_input.Filesize, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetFilesize() int { return v.Filesize }

// GetWidth returns Create_directus_files_input.Width, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetWidth() int { return v.Width }

// GetHeight returns Create_directus_files_input.Height, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetHeight() int { return v.Height }

// GetDuration returns Create_directus_files_input.Duration, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetDuration() int { return v.Duration }

// GetEmbed returns Create_directus_files_input.Embed, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetEmbed() string { return v.Embed }

// GetDescription returns Create_directus_files_input.Description, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetDescription() string { return v.Description }

// GetLocation returns Create_directus_files_input.Location, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetLocation() string { return v.Location }

// GetTags returns Create_directus_files_input.Tags, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetTags() string { return v.Tags }

// GetMetadata returns Create_directus_files_input.Metadata, and is useful for accessing the field via an interface.
func (v *Create_directus_files_input) GetMetadata() string { return v.Metadata }

type Create_directus_folders_input struct {
	Id     string                         `json:"id"`
	Name   string                         `json:"name"`
	Parent *Create_directus_folders_input `json:"parent,omitempty"`
}

// GetId returns Create_directus_folders_input.Id, and is useful for accessing the field via an interface.
func (v *Create_directus_folders_input) GetId() string { return v.Id }

// GetName returns Create_directus_folders_input.Name, and is useful for accessing the field via an interface.
func (v *Create_directus_folders_input) GetName() string { return v.Name }

// GetParent returns Create_directus_folders_input.Parent, and is useful for accessing the field via an interface.
func (v *Create_directus_folders_input) GetParent() *Create_directus_folders_input { return v.Parent }

type Create_directus_roles_input struct {
	Id           string                         `json:"id"`
	Name         string                         `json:"name"`
	Icon         string                         `json:"icon"`
	Description  string                         `json:"description"`
	Ip_access    []string                       `json:"ip_access"`
	Enforce_tfa  bool                           `json:"enforce_tfa"`
	Admin_access bool                           `json:"admin_access"`
	App_access   bool                           `json:"app_access"`
	Users        []*Create_directus_users_input `json:"users,omitempty"`
}

// GetId returns Create_directus_roles_input.Id, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetId() string { return v.Id }

// GetName returns Create_directus_roles_input.Name, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetName() string { return v.Name }

// GetIcon returns Create_directus_roles_input.Icon, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetIcon() string { return v.Icon }

// GetDescription returns Create_directus_roles_input.Description, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetDescription() string { return v.Description }

// GetIp_access returns Create_directus_roles_input.Ip_access, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetIp_access() []string { return v.Ip_access }

// GetEnforce_tfa returns Create_directus_roles_input.Enforce_tfa, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetEnforce_tfa() bool { return v.Enforce_tfa }

// GetAdmin_access returns Create_directus_roles_input.Admin_access, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetAdmin_access() bool { return v.Admin_access }

// GetApp_access returns Create_directus_roles_input.App_access, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetApp_access() bool { return v.App_access }

// GetUsers returns Create_directus_roles_input.Users, and is useful for accessing the field via an interface.
func (v *Create_directus_roles_input) GetUsers() []*Create_directus_users_input { return v.Users }

type Create_directus_users_input struct {
	Id                  string                       `json:"id"`
	First_name          string                       `json:"first_name"`
	Last_name           string                       `json:"last_name"`
	Email               string                       `json:"email"`
	Password            string                       `json:"password"`
	Location            string                       `json:"location"`
	Title               string                       `json:"title"`
	Description         string                       `json:"description"`
	Tags                string                       `json:"tags"`
	Avatar              *Create_directus_files_input `json:"avatar,omitempty"`
	Language            string                       `json:"language"`
	Theme               string                       `json:"theme"`
	Tfa_secret          string                       `json:"tfa_secret"`
	Status              string                       `json:"status"`
	Role                *Create_directus_roles_input `json:"role,omitempty"`
	Token               string                       `json:"token"`
	Last_access         time.Time                    `json:"last_access"`
	Last_access_func    *Datetime_functionsInput     `json:"last_access_func,omitempty"`
	Last_page           string                       `json:"last_page"`
	Provider            string                       `json:"provider"`
	External_identifier string                       `json:"external_identifier"`
	Auth_data           string                       `json:"auth_data"`
	Email_notifications bool                         `json:"email_notifications"`
}

// GetId returns Create_directus_users_input.Id, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetId() string { return v.Id }

// GetFirst_name returns Create_directus_users_input.First_name, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetFirst_name() string { return v.First_name }

// GetLast_name returns Create_directus_users_input.Last_name, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLast_name() string { return v.Last_name }

// GetEmail returns Create_directus_users_input.Email, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetEmail() string { return v.Email }

// GetPassword returns Create_directus_users_input.Password, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetPassword() string { return v.Password }

// GetLocation returns Create_directus_users_input.Location, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLocation() string { return v.Location }

// GetTitle returns Create_directus_users_input.Title, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetTitle() string { return v.Title }

// GetDescription returns Create_directus_users_input.Description, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetDescription() string { return v.Description }

// GetTags returns Create_directus_users_input.Tags, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetTags() string { return v.Tags }

// GetAvatar returns Create_directus_users_input.Avatar, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetAvatar() *Create_directus_files_input { return v.Avatar }

// GetLanguage returns Create_directus_users_input.Language, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLanguage() string { return v.Language }

// GetTheme returns Create_directus_users_input.Theme, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetTheme() string { return v.Theme }

// GetTfa_secret returns Create_directus_users_input.Tfa_secret, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetTfa_secret() string { return v.Tfa_secret }

// GetStatus returns Create_directus_users_input.Status, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetStatus() string { return v.Status }

// GetRole returns Create_directus_users_input.Role, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetRole() *Create_directus_roles_input { return v.Role }

// GetToken returns Create_directus_users_input.Token, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetToken() string { return v.Token }

// GetLast_access returns Create_directus_users_input.Last_access, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLast_access() time.Time { return v.Last_access }

// GetLast_access_func returns Create_directus_users_input.Last_access_func, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLast_access_func() *Datetime_functionsInput {
	return v.Last_access_func
}

// GetLast_page returns Create_directus_users_input.Last_page, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetLast_page() string { return v.Last_page }

// GetProvider returns Create_directus_users_input.Provider, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetProvider() string { return v.Provider }

// GetExternal_identifier returns Create_directus_users_input.External_identifier, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetExternal_identifier() string { return v.External_identifier }

// GetAuth_data returns Create_directus_users_input.Auth_data, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetAuth_data() string { return v.Auth_data }

// GetEmail_notifications returns Create_directus_users_input.Email_notifications, and is useful for accessing the field via an interface.
func (v *Create_directus_users_input) GetEmail_notifications() bool { return v.Email_notifications }

type Datetime_functionsInput struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Week    int `json:"week"`
	Day     int `json:"day"`
	Weekday int `json:"weekday"`
	Hour    int `json:"hour"`
	Minute  int `json:"minute"`
	Second  int `json:"second"`
}

// GetYear returns Datetime_functionsInput.Year, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetYear() int { return v.Year }

// GetMonth returns Datetime_functionsInput.Month, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetMonth() int { return v.Month }

// GetWeek returns Datetime_functionsInput.Week, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetWeek() int { return v.Week }

// GetDay returns Datetime_functionsInput.Day, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetDay() int { return v.Day }

// GetWeekday returns Datetime_functionsInput.Weekday, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetWeekday() int { return v.Weekday }

// GetHour returns Datetime_functionsInput.Hour, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetHour() int { return v.Hour }

// GetMinute returns Datetime_functionsInput.Minute, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetMinute() int { return v.Minute }

// GetSecond returns Datetime_functionsInput.Second, and is useful for accessing the field via an interface.
func (v *Datetime_functionsInput) GetSecond() int { return v.Second }

// __CreateAssetInput is used internally by genqlient
type __CreateAssetInput struct {
	Data *Create_assets_input `json:"data,omitempty"`
}

// GetData returns __CreateAssetInput.Data, and is useful for accessing the field via an interface.
func (v *__CreateAssetInput) GetData() *Create_assets_input { return v.Data }

func CreateAsset(
	ctx context.Context,
	client graphql.Client,
	data *Create_assets_input,
) (*CreateAssetResponse, error) {
	__input := __CreateAssetInput{
		Data: data,
	}
	var err error

	var retval CreateAssetResponse
	err = client.MakeRequest(
		ctx,
		"CreateAsset",
		`
mutation CreateAsset ($data: create_assets_input!) {
	create_assets_item(data: $data) {
		id
		name
		duration
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
