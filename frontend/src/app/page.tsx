'use client'

import * as React from 'react'
import { useState } from 'react'
import { FileUploader } from '../components/file-uploader'
import { QueryResults } from '../components/query-results'

export default function Home() {
  const [results, setResults] = useState<any>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const handleFileUpload = async (file: File) => {
    setLoading(true)
    setError(null)
    
    try {
      const formData = new FormData()
      formData.append('transcript', file)
      
      // Connect to the Go backend API
      const response = await fetch('http://localhost:8080/api/analyze', {
        method: 'POST',
        body: formData,
      })
      
      if (!response.ok) {
        throw new Error(`Error: ${response.status} - ${await response.text()}`)
      }
      
      const data = await response.json()
      setResults(data)
    } catch (err) {
      console.error('Error during API call:', err)
      setError(err instanceof Error ? err.message : 'An error occurred connecting to the backend')
      setResults(null)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="container mx-auto py-6 px-4">
          <h1 className="text-3xl font-bold text-gray-900">Sales Quotation AI</h1>
          <p className="text-gray-500 mt-1">Upload your transcript to analyze sales data</p>
        </div>
      </header>
      
      <main className="container mx-auto py-10 px-4 max-w-6xl">
        <div className="mb-10">
          <FileUploader onFileUpload={handleFileUpload} loading={loading} />
        </div>
        
        {error && (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-6">
            {error}
          </div>
        )}
        
        {results && <QueryResults results={results} />}
      </main>
      
      <footer className="bg-white border-t mt-auto">
        <div className="container mx-auto py-4 px-4 text-center text-gray-500 text-sm">
          © {new Date().getFullYear()} Sales Quotation AI. All rights reserved.
        </div>
      </footer>
    </div>
  )
}
