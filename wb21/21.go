// Адаптер(Adapter) - структурный шаблон проектирования,
// предназначенный для организации использования функции объекта,
// недоступно для модификаций, через специально созданный интерфейс

// Проблема:
// У нас уже есть конкретная структура и нужно,
// чтобы эта структура реализовывала определенный интерфейс,
// при этом структуру менять нельзя

package main

import "fmt"

// внешний сервис, который работает в формате XML. Мы отправляем данные в формате XML
// он всегда неизменный
type AnalyticalDataService interface {
	SendXmlData()
}

// для того, чтобы работать с сервисом аналитики, то нам нужен XML документ
type XmlDocument struct{}

// реализуем интерфейс AnalyticalDataService, который он использует
func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка xml документа")
}

func main() {
}

// все объекты, которые мы ипользуем имеют json формат
type JsonDocument struct{}

// Функция, которая конвертирует xml объект в json структуру
func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

// АДАПТЕР
// мы подставляем структуру JsonDocument, который наследует наш документ
// тем самым мы избегаем связанности с конкретным сервисом (AnalyticalDataService)
type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

// используем наш адаптер и обращаемся к нашему документу и уже у этого документы мы вызываем ConvertToXml
// после конвертации мы говорим, что отправляем наши данные
func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Отправка xml данных в сервис аналитики")
}
