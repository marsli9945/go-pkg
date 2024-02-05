package util

const (
	FEISHU_WEBHOOK_AUTO_CLICK = "https://open.feishu.cn/open-apis/bot/v2/hook/e84400bd-2821-4005-84a2-96b71c83b539"
	FEISHU_WEBHOOK_MARS_ALARM = "https://open.feishu.cn/open-apis/bot/v2/hook/5c7e7a67-6c38-4b34-aa24-49cf4d9c442a"
	Signature                 = "【线上告警】"
)

func FeiShuMessage(message string, webhookAddress string) error {
	_, err := PostAndHeader(
		webhookAddress,
		map[string]any{
			"msg_type": "text",
			"content": map[string]string{
				"text": message,
			},
		},
		map[string]string{
			"Content-Type": "application/json",
		},
	)
	if err != nil {
		return err
	}
	return nil
}
