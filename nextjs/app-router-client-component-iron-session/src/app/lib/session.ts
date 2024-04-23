import { SessionOptions } from 'iron-session';

export interface SessionData {
	email: string;
	username: string;
	isLoggedIn: boolean;
}

export const defaultSession: SessionData = {
	email: '',
	username: '',
	isLoggedIn: false,
};

export const sessionOptions: SessionOptions = {
	password: 'r4nd0mp4ssw0rdaksjhdasuk918ajksdhka',
	cookieName: 'playground-iron-session',
	cookieOptions: {
		// secure only works in `https` environments
		// if your localhost is not on `https`, then use: `secure: process.env.NODE_ENV === "production"`
		secure: true,
	},
};

export function sleep(ms: number) {
	return new Promise((resolve) => setTimeout(resolve, ms));
}
