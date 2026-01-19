# GoCalc (Fyne Edition)

A production-grade, cross-platform calculator application built with **Golang** and **Fyne v2**.

This project serves as a reference implementation for building GUI applications in Go that adhere to **industry standards**. It demonstrates strict separation of concerns, decoupling business logic from the presentation layer, and implementing robust unit testing.

## 🏗 Architecture

Unlike typical tutorial code, this project avoids "spaghetti code" by strictly isolating the math engine from the UI pixels.

```text
go-calculator/
├── cmd/
│   └── app/          # Main entry point (Wiring)
├── internal/
│   ├── calculator/   # Pure Business Logic (No UI code)
│   └── ui/           # Fyne Presentation Layer
└── go.mod

```

* **`internal/calculator`**: Contains the `Engine` struct. It handles state, parsing, and arithmetic. It has zero dependencies on Fyne.
* **`internal/ui`**: Handles the window, buttons, and layout. It talks to the `Engine` via methods, not direct field access.

## 🚀 Getting Started

### Prerequisites

* Go 1.25 or higher
* C compiler (GCC/MinGW) - required by Fyne for CGO bindings.

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-calculator.git
cd go-calculator

```


2. Install dependencies:
```bash
go mod tidy

```


3. Run the application:
```bash
go run cmd/app/main.go

```



## 📦 Building for Production

To create a native executable (specifically for Windows) without the debug console popping up, use the linker flags:

### Manual Build

```powershell
go build -ldflags "-H=windowsgui" -o calculator.exe cmd/app/main.go

```

### Using Fyne Tool (Recommended)

This method bundles icons and metadata automatically.

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os windows -icon Icon.png --src cmd/app

```

## 🧪 Testing

The business logic is fully covered by unit tests. Because the logic is decoupled from the GUI, we can test mathematical edge cases without launching a window.

**Run all tests:**

```bash
go test -v ./internal/calculator

```

**Check coverage:**

```bash
go test -cover ./internal/calculator

```

## 🛠 Tech Stack

* **Language**: Go (Golang)
* **GUI Toolkit**: [Fyne v2](https://fyne.io/) - Native rendering (OpenGL), not a webview.
* **Architecture**: Modular / Clean Architecture.

## 📝 License

Working on it.