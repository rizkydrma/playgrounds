import { FormEvent, useState } from "react";
import { Socket } from "socket.io-client";

const ChatFooter = ({ socket }: { socket: Socket }) => {
	const [message, setMessage] = useState("");

	const handleSendMessage = (e: FormEvent) => {
		e.preventDefault();
		if (message.trim() && localStorage.getItem("userName")) {
			socket.emit("message", {
				text: message,
				name: localStorage.getItem("userName"),
				id: `${socket.id}${Math.random()}`,
				socketID: socket.id,
			});
		}
		setMessage("");
	};
	return (
		<div className="chat__footer">
			<form className="form" onSubmit={handleSendMessage}>
				<input
					type="text"
					placeholder="Write message"
					className="message"
					value={message}
					onChange={(e) => setMessage(e.target.value)}
				/>
				<button className="sendBtn">SEND</button>
			</form>
		</div>
	);
};

export default ChatFooter;
