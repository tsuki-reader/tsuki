import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    /* config options here */
    output: "export",
    allowedDevOrigins: ["wails.localhost"],
    trailingSlash: true,
};

export default nextConfig;
