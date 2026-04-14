package tuairisc

import (
	"NewsChannel/news"
)

func (a *tuairisc) GetArticles() ([]news.Article, error) {
	var articles []news.Article

	temp, err := a.GetNationalArticles()
	if err != nil {
		return nil, err
	}
	articles = append(articles, temp...)

	return articles, nil
}

// Need to find out how articles are actually integrated within the channel

func (a *tuairisc) GetNationalArticles() ([]news.Article, error) {
	return a.getArticles("https://www.tuairisc.ie/rss", news.NationalNews)
}
