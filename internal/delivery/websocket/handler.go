package websocket

import (
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "sync"
)

type Handler struct {
    clients    map[*websocket.Conn]bool
    broadcast  chan []byte
    register   chan *websocket.Conn
    unregister chan *websocket.Conn
    mutex      sync.Mutex
}

func NewHandler() *Handler {
    return &Handler{
        clients:    make(map[*websocket.Conn]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *websocket.Conn),
        unregister: make(chan *websocket.Conn),
    }
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Configure this appropriately in production
    },
}

func (h *Handler) Connect(c echo.Context) error {
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    
    h.register <- ws
    
    // Handle client messages
    go h.handleClient(ws)
    return nil
}