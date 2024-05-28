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
	response, err := s.llm.Call(context.Background(), "Generate 30 MCQs about Mitosis and Meiosis (Evenly Mixed), with five options for each (Attach the answer index, starting from 0) [Use Json list format, with fields: question, choices, answer]")
	if err != nil {
		return nil, err
	}

	var questions []*question
	if json.Unmarshal([]byte(strings.ReplaceAll(strings.ReplaceAll(response, "```", ""), "json", "")), &questions) != nil {
		return nil, err
	}

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
