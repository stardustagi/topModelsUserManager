// export const baseUrlApi = (url: string) => `http://127.0.0.1:8080/api${url}`;
// // export const baseUrlApi = (url: string) =>
// //   `http://login.dev.topapi.ai/api${url}`;

// // export const baseUrlApi = (url: string) => `https://login.topapi.ai/api${url}`;

// export const statsUrlApi = (url: string) => `http://127.0.0.1:8888/api${url}`;

export function baseUrlApi(
  url: string,
  service: "main" | "stats" | "image" | "video" | "node" = "main"
) {
  const baseMap = {
    main: import.meta.env.VITE_API_BASE_URL,
    stats: import.meta.env.VITE_STATS_URL,
    image: import.meta.env.VITE_IMAGE_API,
    video: import.meta.env.VITE_VIDEO_API
  };
  console.log("basemap service === ", baseMap[service]);
  return `${baseMap[service]}${url}`;
}
