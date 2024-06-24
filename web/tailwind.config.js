/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        // svelte-like theme
        primary: {
          DEFAULT: "#ff3e00",
          foreground: "#fff",
        },
        background: "#f8f9fa",
        foreground: "#212529",
        secondary: "#dee2e6",
        border: "#dee2e6",
      },
      borderRadius: {
        sm: "0.125rem",
        DEFAULT: "0.25rem",
        md: "0.375rem",
        lg: "0.5rem",
        xl: "0.75rem",
        "2xl": "1rem",
      },
    },
  },
  plugins: [],
};
