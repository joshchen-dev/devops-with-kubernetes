import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  images: {
    dangerouslyAllowLocalIP: true,
    minimumCacheTTL: 0,
    remotePatterns: [
      {
        protocol: "http",
        hostname: "localhost"
      }
    ],
  },
};

export default nextConfig;
