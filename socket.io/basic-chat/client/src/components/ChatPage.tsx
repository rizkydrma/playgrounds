import { Socket } from "socket.io-client";
import ChatBar from "./ChatBar";
import ChatBody from "./ChatBody";
import ChatFooter from "./ChatFooter";
import { useEffect, useState } from "react";

const ChatPage = ({ socket }: { socket: Socket }) => {
	const [messages, setMessages] = useState([]);

	useEffect(() => {
		socket.on("messageResponse", (data) => setMessages([...messages, data]));
	}, [socket, messages]);

	return (
		<div className="chat">
			<ChatBar socket={socket} />
			<div className="chat__main">
				<ChatBody messages={messages} />
				<ChatFooter socket={socket} />
			</div>
		</div>
	);
};

export default ChatPage;
