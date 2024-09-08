import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { SERVER_URL } from "$env/static/private";
import z from "zod";

export type Anime = {
    data: {
        mal_id: number;
        title: string;
        url: string;
        synopsis: string;
        images: {
            webp: {
                image_url: string;
                small_image_url: string;
                large_image_url: string;
            };
        };
    };
};

export const load = (async ({ params }) => {
    const id = params.anime_id;
    const anime = await api<Anime>(`https://api.jikan.moe/v4/anime/${id}`);
    if (!anime.success) {
        console.error("Failed to fetch anime", anime.error);
        throw error(500, "Failed to fetch anime");
    }
    console.debug("anime", anime.data);
    return {
        anime: anime.data.data,
    };
}) satisfies PageServerLoad;

export const actions = {
    addToFavorites: async ({ request }) => {
        const form = await request.formData();
        const requestBody = await z
            .object({
                mal_id: z.string(),
                title: z.string(),
                image: z
                    .string()
                    .startsWith("https://cdn.myanimelist.net/images"),
            })
            .safeParseAsync({
                mal_id: form.get("mal_id"),
                title: form.get("title"),
                image: form.get("image"),
            });

        if (requestBody.success) {
            const response = await api(`${SERVER_URL}/favorites`, {
                method: "POST",
                body: requestBody.data,
            });

            if (!response.success) {
                return { ...response };
            }

            return { success: response.success };
        } else {
            return { success: false };
        }
    },

    removeFromFavorites: async ({ request }) => {
        const form = await request.formData();

        const mal_id = form.get("mal_id") as unknown as number;

        const response = await api(`${SERVER_URL}/favorites/${mal_id}`, {
            method: "DELETE",
        });

        if (!response.success) {
            return { ...response };
        }

        return { success: response.success };
    },
} satisfies Actions;
