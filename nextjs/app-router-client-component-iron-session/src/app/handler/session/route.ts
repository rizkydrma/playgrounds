import { NextRequest } from 'next/server';
import { cookies } from 'next/headers';
import { getIronSession } from 'iron-session';
import { defaultSession, sessionOptions, sleep, SessionData } from '../../lib/session';

// login
export async function POST(request: NextRequest) {
	const session = await getIronSession<SessionData>(cookies(), sessionOptions);

	const { username = 'No username', email } = (await request.json()) as {
		username: string;
		email: string;
	};

	session.isLoggedIn = true;
	session.username = username;
	session.email = email;
	await session.save();

	// simulate looking up the user in db
	// await sleep(250);

	return Response.json(session);
}

// read session
export async function GET() {
	const session = await getIronSession<SessionData>(cookies(), sessionOptions);

	// simulate looking up the user in db
	// await sleep(250);

	if (session.isLoggedIn !== true) {
		return Response.json(defaultSession);
	}

	return Response.json(session);
}

// logout
export async function DELETE() {
	const session = await getIronSession<SessionData>(cookies(), sessionOptions);

	session.destroy();

	return Response.json(defaultSession);
}
