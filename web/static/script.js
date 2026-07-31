const socketInput = document.querySelector("#socket-input")
const sendButton = document.querySelector("#send")
const socketResponse = document.querySelector("#socket-response")
const getInfoButton = document.querySelector("#get-info")
const getStatusButton = document.querySelector("#get-status")
const fader1 = document.querySelector("#fader-1")

sendButton.addEventListener("click", e => {
    if (socket.readyState == socket.OPEN) {
        const value = socketInput.value
        socket.send(value)
    }
})

getInfoButton.addEventListener("click", e => {
    if (socket.readyState == socket.OPEN) {
        message = {
            op: "GET_OSC",
            key: "/info",
            value: "string value",
        }

        socket.send(JSON.stringify(message))
    }
})

getStatusButton.addEventListener("click", e => {
    if (socket.readyState == socket.OPEN) {
        socket.send("get-status")
    }
})

function fader1Changed() {
    message = {
        op: "SET_OSC",
        key: "/ch/01/mix/fader",
        value: fader1.value / 1024,
    }

    socket.send(JSON.stringify(message))
}

// =======================

const networkIcon = document.querySelector("#network-icon")
const socketStatus = document.querySelector("#socket-status")
socketStatus.textContent = "Connected"
const mixerStatus = document.querySelector("#mixer-status")
const mixerIpInput = document.querySelector("#mixer-ip")
const mixerPortInput = document.querySelector("#mixer-port")
const connectButton = document.querySelector("#mixer-connect")

let awaitCloseConnectionPopup = false
let refreshSocketTimeout
let refreshAttempts = 0
const MAX_REFRESH_ATTEMPTS = 10

const host = window.location.host
let socket

const url = new URL(window.location.href)

const tabElements = document.querySelectorAll(".tab")
const tabs = Array.from(tabElements).reduce((tabs, tab) => ({
    ...tabs,
    [tab.id]: tab,
}), {})
const pageElements = document.querySelectorAll(".page")
const pages = Array.from(pageElements).reduce((pages, page) => ({
    ...pages,
    [page.id]: page,
}), {})

let currentTab = url.searchParams.get("page")
if (!currentTab) {
    currentTab = Object.values(tabs)[0].id
    url.searchParams.set("page", currentTab)
    window.history.replaceState({}, "", url)
}

Object.values(tabs).forEach((tab) => {
    const id = tab.id
    if (id == currentTab){
        tab.classList.add("selected")
    }
    tab.addEventListener("click", e => {
        tabs[currentTab].classList.remove("selected")
        pages[currentTab].hidden = true
        currentTab = id
        tab.classList.add("selected")
        pages[currentTab].hidden = false

        url.searchParams.set("page", id)
        window.history.replaceState({}, "", url)
    })
})

Object.values(pages).forEach((page) => {
    if (page.id != currentTab) {
        page.hidden = true
    }
})


const popupElements = document.querySelectorAll(".popup")
const popups = Array.from(popupElements).reduce((popups, popup) => ({
    ...popups,
    [popup.id]: popup
}), {})

Object.values(popups).forEach(popup => {
    popup.hidden = true;

    popup.addEventListener("click", e => {
        popup.hidden = true
    })
    popup.querySelector(".popup-content").addEventListener("click", e => {
        e.stopPropagation()
    })
})

function onSocketOpen(event) {
    networkIcon.classList.add("found")
    socketStatus.textContent = "Connected"
    socketStatus.classList.add("connected")

    console.log("Connected to the server!")

    message = {
        op: "GET",
        key: "sync"
    }
    socket.send(JSON.stringify(message))

    refreshAttempts = 0
}

function onSocketMessage(event) {
    message = JSON.parse(event.data)
    if (message.op === "SET_OSC" && message.key === "/ch/01/mix/fader") {
        fader1.value = message.value * 1024
    }
    
    if (message.op === "SET") {
        if (message.key === "mixer-address") {
            mixerIpInput.value = message.value.ip
            mixerPortInput.value = message.value.port
        }
    }

    if (message.op === "SET" && message.key === "status") {
        if (message.value) {
            networkIcon.classList.add("connected")
            mixerStatus.textContent = "Connected"
            mixerStatus.classList.add("connected")
            if (awaitCloseConnectionPopup) {
                popups.connectionModal.hidden = true
                awaitCloseConnectionPopup = false
            }
        } else {
            networkIcon.classList.remove("connected")
            mixerStatus.textContent = "Disconnected"
            mixerStatus.classList.remove("connected")
            popups.connectionModal.hidden = false
        }
    }

    socketResponse.textContent = event.data
}

function onSocketClose(event) {
    networkIcon.classList.remove("connected")
    networkIcon.classList.remove("found")
    socketStatus.textContent = "Disconnected"
    socketStatus.classList.remove("connected")
    mixerStatus.textContent = "Disconnected"
    mixerStatus.classList.remove("connected")

    console.log("Socket closed.")

    refreshAttempts++
    if (refreshAttempts <= MAX_REFRESH_ATTEMPTS) {
        refreshSocketTimeout = setTimeout(refreshSocket, 5 * 1000)
    }
}

function refreshSocket() {
    console.log("Refreshing socket")
    if (socket && socket.readyState == socket.OPEN) {
        return
    }
    clearTimeout(refreshSocketTimeout)

    socket = new WebSocket(`ws://${host}/ws`)

    socket.onopen = onSocketOpen
    socket.onmessage = onSocketMessage
    socket.onclose = onSocketClose
}

refreshSocketTimeout = setTimeout(refreshSocket, 50)

networkIcon.addEventListener("click", e => {
    popups.connectionModal.hidden = false

    if (socket.readyState != socket.OPEN) {
        refreshAttempts = 0
        clearTimeout(refreshSocketTimeout)
        refreshSocket()
        return
    }

    message = {
        op: "GET",
        key: "mixer-address"
    }
    socket.send(JSON.stringify(message))
})

connectButton.addEventListener("click", e => {
    if (socket.readyState != socket.OPEN) {
        return
    }

    awaitCloseConnectionPopup = true

    message = {
        op: "SET",
        key: "mixer-address",
        value: {
            ip: mixerIpInput.value,
            port: parseInt(mixerPortInput.value),
        },
    }
    socket.send(JSON.stringify(message))
})


