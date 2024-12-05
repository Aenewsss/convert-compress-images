package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import necessário para suporte a JPEG
	_ "image/png"  // Import necessário para suporte a PNG
	"os"
	"strings"

	"github.com/chai2010/webp"
)

func main() {
	// Caminho para a imagem de entrada
	inputPath := "image-site2.png"
	outputPath := "output.webp"

	if strings.Contains(inputPath, ".webp") {
		fmt.Println("Imagem já está no formato correto")
		return
	}

	err := convertToWebP(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Erro ao converter imagem: %v\n", err)
		return
	}

	fmt.Printf("Imagem convertida com sucesso: %s\n", outputPath)
}

// convertToWebP converte uma imagem de qualquer formato suportado para WebP
func convertToWebP(inputPath, outputPath string) error {
	// Abrir o arquivo de entrada
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("erro ao abrir a imagem: %w", err)
	}
	defer file.Close()

	// Detectar o formato da imagem
	imageData, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("erro ao decodificar a imagem: %w", err)
	}

	fmt.Printf("Formato detectado: %s\n", format)

	// Criar o arquivo de saída
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo de saída: %w", err)
	}
	defer outputFile.Close()

	// Codificar a imagem no formato WebP
	err = webp.Encode(outputFile, imageData, &webp.Options{Quality: 10})
	if err != nil {
		return fmt.Errorf("erro ao codificar a imagem para WebP: %w", err)
	}

	return nil
}
