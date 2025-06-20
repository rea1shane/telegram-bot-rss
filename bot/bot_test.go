package bot

import (
	"os"
	"testing"
)

var (
	botToken = os.Getenv("TELEGRAM_BOT_KEY")
	chatId   = os.Getenv("TELEGRAM_CHAT_ID")
)

func TestBot_Send(t *testing.T) {
	bot, err := New(botToken, chatId)
	if err != nil {
		t.Fatalf("failed to new bot: %v", err)
	}

	err = bot.Send(
		"阮一峰的网络日志",
		"科技爱好者周刊（第 349 期）：神经网络算法的发明者",
		"http://www.ruanyifeng.com/blog/2025/05/weekly-issue-349.html",
		"https://cdn.beekka.com/blogimg/asset/202505/bg2025052105.webp",
		[]string{"Weekly", "Blog"},
	)
	if err != nil {
		t.Fatalf("failed to send notification: %v", err)
	}
}
