function changeMessage(){
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;

    const alternateMessages = [
        "¡Hola de nuevo!",
        "¡Bienvenido de vuelta!",
        "¡Que tengas un maravilloso día!",
    ]

    let newMessage;

    do {
        newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)]
    } while (newMessage === currentMessage);

    messageElement.textContent = newMessage;
}