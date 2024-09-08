import { SERVER_URL } from "$env/static/private";

import api from "$lib/server/api";
import type { LayoutServerLoad } from "./$types";

type Favorite = {
    mal_id: number;
    title: string;
    image: string;
};

export const load = (async () => {
    const go_server_favorites = await api<Array<Favorite>>(
        SERVER_URL + "/favorites",
    );
    if (!go_server_favorites.success) {
        console.error(
            "Failed to load favorites from server",
            go_server_favorites.error,
        );

        return new Error("Failed to load favorites from server");
    } else {
        return {
            favorites: go_server_favorites.data,
        };
    }
}) satisfies LayoutServerLoad;
