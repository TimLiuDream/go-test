package matchers

import (
	"encoding/xml"
	"GolangTraining/simple1/search"
	"errors"
	"net/http"
	"fmt"
	"log"
	"regexp"
)

type (
	//item根据item字段的标签，将定义的字段与rss文档的字段关联起来
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
	//image根据image字段的标签，将定义的字段与rss文档的字段关联起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	//channel根据channel字段的标签，将定义的字段与rss文档的字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}
	//rssDocument定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

//rssMatcher实现了Matcher接口
type rssMatcher struct{}

//初始化注册匹配器
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

//retieve发送http get请求获取rss数据源并解码
func (m rssMatcher) retieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("no rss feed uri provided")
	}

	//从网络获得rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	//一旦函数返回，关闭返回的响应链接
	defer resp.Body.Close()

	//检查状态码是否200，这样就知道是不是收到了正确的数据
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http response error %d\n", resp.StatusCode)
	}

	//将rss数据源文档解码到我们定义的结构类型里，不需要检查错误，调用者会做这件事
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

//search在文档中查找特定的搜索项
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("search feed type[%s] site[%s] for uri[%s]\n", feed.Type, feed.Name, feed.URI)

	//获取想要搜索的数据
	document, err := m.retieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		//检查标题是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		//如果找到匹配项，将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		//检查 描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
