'use client';

import { useRouter } from 'next/navigation';
import { defaultSession } from '../lib/session';
import useSession from '../lib/use-session';

export default function LoginFormPage() {
	const { session, isLoading } = useSession();

	if (isLoading) {
		return <p className="text-lg">Loading...</p>;
	}

	if (!session.isLoggedIn) return <LoginForm />;

	return (
		<>
			<p className="text-lg">
				Logged in user: <strong>{session.username}</strong>
			</p>
			<LogoutButton />
		</>
	);
}

function LoginForm() {
	const { login } = useSession();
	const navigation = useRouter();

	return (
		<div className="w-full min-h-[100vh] grid place-items-center">
			<div className="min-w-[400px] p-4 bg-gray-200 border rounded-md">
				<form
					className="flex flex-col gap-2"
					method="POST"
					onSubmit={async (event) => {
						try {
							event.preventDefault();
							const formData = new FormData(event.currentTarget);
							const username = formData.get('username') as string;
							const email = formData.get('email') as string;

							await login({ username, email });
						} catch (error) {
							console.log({ error });
						}
					}}
				>
					<div>
						<label htmlFor="username">Username</label>
						<input type="text" name="username" className="w-full p-2" placeholder="username...." />
					</div>
					<div>
						<label htmlFor="email">email</label>
						<input type="text" name="email" className="w-full p-2" placeholder="email...." />
					</div>

					<button type="submit" className="bg-blue-600 text-white rounded-sm shadow w-full px-3 py-1">
						Login
					</button>
				</form>
			</div>
		</div>
	);
}

function LogoutButton() {
	const { logout } = useSession();

	return (
		<p>
			<a
				onClick={(event) => {
					event.preventDefault();
					logout(null, {
						optimisticData: defaultSession,
					});
				}}
			>
				Logout
			</a>
		</p>
	);
}
