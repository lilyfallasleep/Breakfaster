package mybot

import (
	"breakfaster/repository/schema"
	"breakfaster/service/constant"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

const welcomeCard string = `{
  	"type": "bubble",
  	"body": {
    	"type": "box",
    	"layout": "vertical",
    	"contents": [
    		{
      			"type": "text",
      			"text": "LINE ",
     			 "weight": "bold",
      			"color": "#1DB446",
     			 "size": "sm"
    		},
			{
				"type": "text",
				"text": "Breakfaster",
				"weight": "bold",
				"size": "xxl",
				"margin": "md"
			},
			{
				"type": "text",
				"text": "嗨嗨～歡迎來到Breakfaster早餐點餐系統！",
				"size": "xs",
				"color": "#1DB446",
				"wrap": true,
				"margin": "xs"
			},
			{
				"type": "text",
				"text": "以下點餐規則請詳細閱讀",
				"size": "xs",
				"color": "#3B3C40",
				"wrap": true,
				"margin": "xs"
			},
			{
				"type": "text",
				"text": "跟著我們一同快樂享用早餐吧～",
				"size": "xs",
				"color": "#3B3C40",
				"wrap": true,
				"margin": "xs"
			},
			{
				"type": "separator",
				"margin": "sm",
				"color": "#313540"
			},
   			{
      			"type": "box",
      			"layout": "vertical",
      			"margin": "xxl",
      			"spacing": "md",
      			"contents": [
      				{
        				"type": "text",
        				"text": "點餐規則",
						"margin": "none",
						"size": "xs",
						"color": "#1DB446"
      				},
					{
						"type": "text",
						"text": "1. 每週一開放點餐，每週五中午11:59收單",
						"margin": "xs",
						"size": "xs"
					},
					{
						"type": "text",
						"text": "2. 一次點一週的份量，如果當天「不用」早",
						"margin": "xs",
						"size": "xs"
					},
					{
						"type": "text",
						"text": "餐請點選「略過」",
						"margin": "none",
						"size": "xs"
					},
					{
						"type": "text",
						"text": "3. 收單前皆可更改/取消訂單，請點擊「來點",
						"margin": "xs",
						"size": "xs"
					},
					{
						"type": "text",
						"text": "早餐吧」重新填寫表單即可，無次數限制",
						"margin": "none",
						"size": "xs"
					},
					{
						"type": "text",
						"text": "4. 請當個乖寶寶～點餐後務必要取餐喔！",
						"margin": "xs",
						"size": "xs"
					}
				]
			},
			{
				"type": "image",
				"url": "https://www.iconfinder.com/data/icons/street-food-and-food-trucker-1/64/hamburger-fast-food-patty-bread-128.png",
				"margin": "xl"
			}
    	]
  	},
	"styles": {
		"body": {
			"backgroundColor": "#F2E6A7"
		},
		"footer": {
			"separator": true
		}
	}
}
`

// DailyOrders is the order template type for confirmation card
type DailyOrders map[string]string

func getNumOrders(orders *[]schema.SelectOrder) int {
	numOrders := 0
	for _, order := range *orders {
		if order.FoodName != constant.DummyFoodName {
			numOrders++
		}
	}
	return numOrders
}

func getConfirmCardDate(t time.Time) string {
	return fmt.Sprintf("%d-%02d/%02d  %s.", t.Year(), t.Month(), t.Day(), t.Weekday().String()[:3])
}

// GetConfirmOrderItems returns the order template for confirmation card
func GetConfirmOrderItems(orders *[]schema.SelectOrder, start, end time.Time) *DailyOrders {
	dailyOrders := make(DailyOrders)

	for !start.After(end) {
		dailyOrders[getConfirmCardDate(start)] = constant.DummyFoodName
		start = start.AddDate(0, 0, 1)
	}
	for _, order := range *orders {
		dailyOrders[getConfirmCardDate(order.Date)] = order.FoodName
	}
	return &dailyOrders
}

// IntPtr is a helper function for using *int values
func IntPtr(v int) *int {
	return &v
}

// NewWelcomeCard is a factory for welcome message
func NewWelcomeCard() linebot.FlexContainer {
	content, _ := linebot.UnmarshalFlexMessageJSON([]byte(welcomeCard))
	return content
}

// NewMenu is a factory for menu message
func NewMenu() linebot.FlexContainer {
	menu := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:        linebot.FlexComponentTypeImage,
			URL:         "https://images.unsplash.com/photo-1493770348161-369560ae357d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1050&q=80",
			Size:        linebot.FlexImageSizeTypeFull,
			AspectRatio: linebot.FlexImageAspectRatioType20to13,
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
			Action:      &linebot.URIAction{URI: OrderPageURI},
		},
		Body: &linebot.BoxComponent{
			Type:    linebot.FlexComponentTypeBox,
			Layout:  linebot.FlexBoxLayoutTypeVertical,
			Spacing: linebot.FlexComponentSpacingTypeMd,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "Welcome to Breakfaster!",
					Size:   "lg",
					Weight: "bold",
				},
				&linebot.ButtonComponent{
					Type: linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{
						Label: "來點早餐吧",
						URI:   OrderPageURI,
					},
					Margin: "xl",
					Style:  "primary",
					Color:  "#05A66B",
				},

				&linebot.ButtonComponent{
					Type: linebot.FlexComponentTypeButton,
					Action: &linebot.PostbackAction{
						Label: "點餐紀錄",
						Data:  "check_order",
					},
					Style: "primary",
					Color: "#05A66B",
				},
				&linebot.ButtonComponent{
					Type: linebot.FlexComponentTypeButton,
					Action: &linebot.PostbackAction{
						Label: "點餐規則",
						Data:  "rule",
					},
					Style: "primary",
					Color: "#8C8C8C",
				},
				&linebot.ButtonComponent{
					Type: linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{
						Label: "問題回報",
						URI:   OrderPageURI + "/report",
					},
					Style: "primary",
					Color: "#8C8C8C",
				},
			},
		},
	}

	return menu
}

// NewConfirmCard is a factory for order confirmation message
func (app *BreakFaster) NewConfirmCard(lineUID string, start, end time.Time) (linebot.FlexContainer, error) {
	orders, err := app.svc.orderRepo.GetOrdersByLineUID(lineUID, start, end)
	if err != nil {
		return &linebot.BubbleContainer{}, err
	}
	numOrders := getNumOrders(orders)
	template := GetConfirmOrderItems(orders, start, end)

	orderComponent := []linebot.FlexComponent{
		&linebot.TextComponent{
			Type:  linebot.FlexComponentTypeText,
			Text:  "總計：" + strconv.Itoa(numOrders) + "份早餐",
			Size:  "sm",
			Color: "#b7b7b7",
		},
	}

	var sortedDate []string
	for date := range *template {
		sortedDate = append(sortedDate, date)
	}
	sort.Strings(sortedDate)

	for _, date := range sortedDate {
		orderComponent = append(orderComponent, &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    date[5:],
					Size:    "md",
					Gravity: "center",
					Flex:    IntPtr(3),
					Color:   "#b7b7b7",
				},

				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    (*template)[date],
					Size:    "md",
					Gravity: "center",
					Flex:    IntPtr(4),
				},
			},
			Spacing:      "lg",
			CornerRadius: "30px",
			Margin:       "xl",
		})
	}

	confirmCard := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Size: "mega",
		Header: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  "訂單記錄",
							Size:  "sm",
							Color: "#ffffff66",
						},
						&linebot.TextComponent{
							Type:   linebot.FlexComponentTypeText,
							Text:   "您的餐點",
							Size:   "xl",
							Color:  "#ffffff",
							Weight: "bold",
							Flex:   IntPtr(4),
						},
					},
				},
			},
			BackgroundColor: "#05A66B",
			Spacing:         "md",
			Height:          "84px",
		},
		Body: &linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: orderComponent,
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type: linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{
						Label: "重新點餐",
						URI:   OrderPageURI,
					},
					Style: "primary",
					Color: "#05A66B",
				},
			},
		},
	}
	return confirmCard, nil
}
