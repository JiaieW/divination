window.addEventListener('DOMContentLoaded', (event) => {
    const chatToggle = document.getElementById('chat-toggle');
    const chatBox = document.getElementById('chat-box');
    const sendButton = document.getElementById('send-button');
    const messageInput = document.getElementById('message-input');

    chatToggle.addEventListener('click', () => {
        if (chatBox.style.display === "none") {
            chatBox.style.display = "block";
            // 发送初始消息
            sendInitialMessage('你好呀');
        } else {
            chatBox.style.display = "none";
        }
    });


    sendButton.addEventListener('click', () => {
        postMessage();
    });

    messageInput.addEventListener('keypress', (event) => {
        if (event.key === 'Enter' && !sendButton.disabled) {
            event.preventDefault();
            postMessage();
        }
    });

    function postMessage() {
        const message = messageInput.value.trim();
        if (message) {
            addMessageToChat('You', message);
            messageInput.value = '';
            sendButton.disabled = true;

            fetch('/message', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ message: message })
            })
            .then(response => response.json())
            .then(data => {
                addMessageToChat('GPT', data.reply);
                sendButton.disabled = false;
            })
            .catch(error => {
                console.error('Error:', error);
                sendButton.disabled = false;
            });
        }
    }

    function addMessageToChat(sender, message) {
        const chat = document.getElementById('chat');
        const messageDiv = document.createElement('div');
        messageDiv.textContent = `${sender}: ${message}`;
        chat.appendChild(messageDiv);
        chat.scrollTop = chat.scrollHeight;
    }

    function sendInitialMessage(message) {
        addMessageToChat('You', message);
        sendButton.disabled = true;

        fetch('/message', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ message: message })
        })
        .then(response => response.json())
        .then(data => {
            addMessageToChat('GPT', data.reply);
            sendButton.disabled = false;
        })
        .catch(error => {
            console.error('Error:', error);
            sendButton.disabled = false;
        });
    }
});
