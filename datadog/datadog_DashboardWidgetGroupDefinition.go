// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetGroupDefinition struct {
	// The layout type of the group. Valid values are `ordered`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#layout_type Dashboard#layout_type}
	LayoutType *string `field:"required" json:"layoutType" yaml:"layoutType"`
	// The background color of the group title, options: `vivid_blue`, `vivid_purple`, `vivid_pink`, `vivid_orange`, `vivid_yellow`, `vivid_green`, `blue`, `purple`, `pink`, `orange`, `yellow`, `green`, `gray` or `white`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#background_color Dashboard#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The image URL to display as a banner for the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#banner_img Dashboard#banner_img}
	BannerImg *string `field:"optional" json:"bannerImg" yaml:"bannerImg"`
	// Whether to show the title or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#show_title Dashboard#show_title}
	ShowTitle interface{} `field:"optional" json:"showTitle" yaml:"showTitle"`
	// The title of the group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#widget Dashboard#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

