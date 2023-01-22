import axios from "./axios";
import type { Album } from "@/models/Album";

export const addNewAlbum = async (form: Album) => {
  try {
    const resp = await axios.post("/albums", {
      title: form.title,
      price: form.price,
      artist: form.artist,
    });
    return {
      data: resp.data,
      status: resp.status,
    };
  } catch (error: any) {
    return {
      data: error.response,
      status: error.response.status,
    };
  }
};
export const getAlbumByTitle = async (title: string) => {
  try {
    const resp = await axios.get(`/album?title=${title}`);
    return {
      data: resp.data.album,
      status: resp.status,
    };
  } catch (error: any) {
    console.log(error.response.status);
    return {
      data: error.response,
      status: error.response.status,
    };
  }
};
