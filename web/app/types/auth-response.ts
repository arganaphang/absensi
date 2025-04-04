import type { User } from "./user";

export type LoginResponse = {
	success: boolean;
	message: string;
	data: {
		token: {
			access_token: string;
			refresh_token: string;
		};
		user: User;
	};
};

export type LogoutResponse = {
	message: string;
};

export type RefreshResponse = {
	success: boolean;
	message: string;
	data: {
		token: {
			access_token: string;
			refresh_token: string;
		};
		user: User;
	};
};

export type AuthResponse = {
	success: boolean;
	message: string;
	data: {
		user: User;
	};
};
