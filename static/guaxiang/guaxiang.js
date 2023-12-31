window.addEventListener('DOMContentLoaded', (event) => {
    const chatToggle = document.getElementById('chat-toggle');
    const chatBox = document.getElementById('chat-box');
    const sendButton = document.getElementById('send-button');
    const messageInput = document.getElementById('message-input');

    chatToggle.addEventListener('click', () => {
        if (chatBox.style.display === "none") {
            chatBox.style.display = "block";
            const benGuaDiv = document.querySelector('.ben_gua');
            const bianGuaDiv = document.querySelector('.bian_gua');
            console.log(benGuaDiv)
            // 提取本卦和变卦的数据
            const guaInfo = {
                benGua: {
                    location: benGuaDiv.getAttribute('data-location'),
                    alias: benGuaDiv.getAttribute('data-alias'),
                    form: benGuaDiv.getAttribute('data-form'),
                    guaci: benGuaDiv.getAttribute('data-guaci'),
                    // 这里添加其他需要的本卦信息
                },
                bianGua: {
                    location: bianGuaDiv.getAttribute('data-location'),
                    alias: bianGuaDiv.getAttribute('data-alias'),
                    form: bianGuaDiv.getAttribute('data-form'),
                    guaci: bianGuaDiv.getAttribute('data-guaci'),
                    // 这里添加其他需要的变卦信息
                },
                // 假设你有一个变化索引的数组，你也可以将其包括进来
                // bianIndexes: 你的变化索引数组
            };
            console.log(guaInfo)
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
        addMessageToChat('您', message);
        sendButton.disabled = true;

        fetch('/message', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ message: message })
        })
        .then(response => response.json())
        .then(data => {
            addMessageToChat('AI周易大师', data.reply);
            sendButton.disabled = false;
        })
        .catch(error => {
            console.error('Error:', error);
            sendButton.disabled = false;
        });
    }
});
