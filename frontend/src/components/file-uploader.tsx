'use client'

import * as React from 'react'
import { useState, useRef } from 'react'
import { Upload, FileText, Loader2 } from 'lucide-react'
import { Card, CardContent } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'
import { cn } from '@/lib/utils'

interface FileUploaderProps {
  onFileUpload: (file: File) => void
  loading: boolean
}

export function FileUploader({ onFileUpload, loading }: FileUploaderProps) {
  const [dragActive, setDragActive] = useState<boolean>(false)
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const [uploadProgress, setUploadProgress] = useState<number>(0)
  const fileInputRef = useRef<HTMLInputElement>(null)

  // Simulate upload progress when loading
  React.useEffect(() => {
    if (loading && uploadProgress < 95) {
      const timer = setTimeout(() => {
        setUploadProgress((prev) => {
          const increment = Math.floor(Math.random() * 10) + 1
          return Math.min(prev + increment, 95)
        })
      }, 500)
      return () => clearTimeout(timer)
    } else if (!loading) {
      setUploadProgress(0)
    }
  }, [loading, uploadProgress])

  const handleDrag = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    
    if (e.type === 'dragenter' || e.type === 'dragover') {
      setDragActive(true)
    } else if (e.type === 'dragleave') {
      setDragActive(false)
    }
  }

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setDragActive(false)
    
    if (e.dataTransfer.files && e.dataTransfer.files[0]) {
      handleFiles(e.dataTransfer.files[0])
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault()
    
    if (e.target.files && e.target.files[0]) {
      handleFiles(e.target.files[0])
    }
  }

  const handleFiles = (file: File) => {
    setSelectedFile(file)
    onFileUpload(file)
  }

  const handleButtonClick = () => {
    fileInputRef.current?.click()
  }

  return (
    <Card className="w-full border-2 shadow-md">
      <CardContent className="p-6">
        <div 
          className={cn(
            "border-2 border-dashed rounded-lg p-10 flex flex-col items-center justify-center",
            dragActive ? "border-primary bg-primary/5" : "border-gray-300",
            loading ? "opacity-70 pointer-events-none" : "hover:border-primary hover:bg-primary/5 cursor-pointer",
            "transition-all duration-200"
          )}
          onDragEnter={handleDrag}
          onDragOver={handleDrag}
          onDragLeave={handleDrag}
          onDrop={handleDrop}
          onClick={handleButtonClick}
        >
          <input
            ref={fileInputRef}
            type="file"
            className="hidden"
            accept=".txt,.doc,.docx,.pdf"
            onChange={handleChange}
            disabled={loading}
          />
          
          {loading ? (
            <div className="flex flex-col items-center space-y-4 w-full">
              <Loader2 className="h-12 w-12 text-primary animate-spin" />
              <p className="text-lg font-medium text-gray-700">Processing transcript...</p>
              <div className="w-full max-w-md mt-2">
                <Progress value={uploadProgress} className="h-2" />
                <p className="text-xs text-gray-500 mt-1 text-right">{uploadProgress}%</p>
              </div>
            </div>
          ) : (
            <>
              <div className="bg-primary/10 p-4 rounded-full mb-4">
                {selectedFile ? (
                  <FileText className="h-10 w-10 text-primary" />
                ) : (
                  <Upload className="h-10 w-10 text-primary" />
                )}
              </div>
              
              <h3 className="text-lg font-medium text-gray-700 mb-1">
                {selectedFile ? selectedFile.name : 'Upload your transcript'}
              </h3>
              
              <p className="text-sm text-gray-500 text-center max-w-sm">
                {selectedFile 
                  ? `${(selectedFile.size / 1024).toFixed(2)} KB · Click to upload a different file` 
                  : 'Drag and drop your file here, or click to browse'}
              </p>
              
              <p className="text-xs text-gray-400 mt-2">
                Supported formats: TXT, DOC, DOCX, PDF
              </p>
            </>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
