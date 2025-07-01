# Sales Quotation AI Frontend

A modern, minimalistic UI for the Sales Quotation AI application built with React, Next.js, and Shadcn UI.

## Features

- Clean, modern UI with minimalistic design
- Transcript upload functionality
- SQL query result visualization
- Responsive design for all devices

## Tech Stack

- **React**: UI library
- **Next.js**: React framework
- **Shadcn UI**: Component library
- **Tailwind CSS**: Utility-first CSS framework
- **TypeScript**: Type-safe JavaScript

## Getting Started

### Prerequisites

- Node.js (v16 or higher)
- npm or yarn

### Installation

1. Clone the repository
2. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
3. Install dependencies:
   ```bash
   npm install --legacy-peer-deps
   ```

### Development

Run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser to see the application.

### Building for Production

Build the application for production:

```bash
npm run build
```

Start the production server:

```bash
npm start
```

## Project Structure

- `/src/app`: Next.js app directory
- `/src/components`: React components
  - `/src/components/ui`: Reusable UI components
- `/src/lib`: Utility functions

## API Integration

The frontend is designed to connect to a Go backend API. The API endpoint for transcript analysis is expected at:

```
http://localhost:8080/api/analyze
```

Currently, the application uses mock data for demonstration purposes.
