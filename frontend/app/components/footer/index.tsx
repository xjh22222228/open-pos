const title = import.meta.env.VITE_TITLE

export default function Footer() {
  return (
    <footer className="text-center text-sm">
      <p>&copy; 2026 {title}. All rights reserved.</p>
    </footer>
  )
}
