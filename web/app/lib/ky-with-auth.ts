import ky, { HTTPError } from "ky";

import type { RefreshResponse } from "~/types/auth-response";
import { ACCESS_TOKEN, REFRESH_TOKEN } from "./constants";

const kyWithAuth = ky.create({
	prefixUrl: import.meta.env.VITE_BASE_URL,
	credentials: "include",
	hooks: {
		beforeRequest: [
			(request) => {
				const accessToken = sessionStorage.getItem(ACCESS_TOKEN);
				if (!accessToken) return;

				request.headers.set("Authorization", `Bearer ${accessToken}`);
			},
		],
		beforeRetry: [
			async ({ request, error, options }) => {
				if (error instanceof HTTPError && error.response.status === 401) {
					const data = await ky
						.post("refresh-token", {
							...options,
							retry: 0,
							json: {
								access_token: sessionStorage.getItem(ACCESS_TOKEN),
								refresh_token: sessionStorage.getItem(REFRESH_TOKEN),
							},
						})
						.json<RefreshResponse>();

					const token = data.data.token;

					sessionStorage.setItem(ACCESS_TOKEN, token.access_token);
					sessionStorage.setItem(REFRESH_TOKEN, token.refresh_token);

					request.headers.set("Authorization", `Bearer ${token.access_token}`);
				}
			},
		],
	},
	retry: {
		limit: 1,
		statusCodes: [401, 403, 500, 504],
	},
});

export { kyWithAuth as ky };
