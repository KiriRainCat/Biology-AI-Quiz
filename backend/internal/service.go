package internal

import (
	"context"
	"encoding/json"
	"log"
	"stp/internal/pkg/config"
	"strings"

	"github.com/tmc/langchaingo/llms/openai"
)

var s *service

type service struct {
	llm *openai.LLM
}

func (s *service) GenerateQuiz() (list []*question, err error) {
	response, err := s.llm.Call(context.Background(), "Generate 16 MCQs about Meiosis, Mitosis, and its checkpoints (Evenly Mixed, no repetition), with 5 options each (Attach answer index, starting from 0) [Use Json list format, with fields: question, choices, answer]")
	if err != nil {
		return nil, err
	}

	// 处理 response
	response = response[strings.Index(response, "[") : strings.LastIndex(response, "]")+1]

	// 将 response 映射到 question 结构
	var questions []*question
	if json.Unmarshal([]byte(response), &questions) != nil {
		return nil, err
	}

	log.Println(response)

	return questions, nil
}

func InitLLM() {
	// 初始化语言模型
	llm, err := openai.New(
		openai.WithBaseURL(config.Config.OpenAI.BaseUrl),
		openai.WithToken(config.Config.OpenAI.Token),
	)
	if err != nil {
		log.Fatalf("LLM 初始化失败: %v", err)
	}

	s = &service{llm: llm}
}
