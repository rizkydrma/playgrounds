import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Socket } from "socket.io-client";

const Home = ({ socket }: { socket: Socket }) => {
	const navigate = useNavigate();
	const [userName, setUserName] = useState("");

	const handleSubmit = (e: FormEvent) => {
		e.preventDefault();
		localStorage.setItem("userName", userName);

		socket.emit("newUser", { userName, socketId: socket.id });
		navigate("/chat");
	};

	return (
		<form className="home__container" onSubmit={handleSubmit}>
			<h2 className="home__header">Sign in to Open Chat</h2>
			<label htmlFor="username">Username</label>
			<input
				type="text"
				minLength={6}
				name="username"
				id="username"
				className="username__input"
				value={userName}
				onChange={(e) => setUserName(e.target.value)}
			/>
			<button className="home__cta">SIGN IN</button>
		</form>
	);
};

export default Home;
