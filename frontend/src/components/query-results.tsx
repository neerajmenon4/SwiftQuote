'use client'

import * as React from 'react'
import { useState } from 'react'
import { ChevronDown, ChevronUp, Database, Code, BarChart } from 'lucide-react'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

interface QueryResultsProps {
  results: {
    query: string
    data: any
  }
}

export function QueryResults({ results }: QueryResultsProps) {
  const [showQuery, setShowQuery] = useState<boolean>(false)
  const [activeView, setActiveView] = useState<'table' | 'chart'>('table')
  
  if (!results || !results.data) {
    return null
  }

  const { query, data } = results
  
  // Ensure data is an array
  const dataArray = Array.isArray(data) ? data : []
  
  // Extract column headers from the first result object
  const columns = dataArray.length > 0 ? Object.keys(dataArray[0]) : []

  return (
    <Card className="shadow-md overflow-hidden">
      <CardHeader className="bg-gray-50 border-b pb-3">
        <div className="flex items-center justify-between">
          <CardTitle className="text-xl">Analysis Results</CardTitle>
          <div className="flex gap-2">
            <Button 
              variant="outline" 
              size="sm"
              className={cn(activeView === 'table' && "bg-primary/10")}
              onClick={() => setActiveView('table')}
            >
              <Database className="h-4 w-4 mr-1" />
              Table
            </Button>
            <Button 
              variant="outline" 
              size="sm"
              className={cn(activeView === 'chart' && "bg-primary/10")}
              onClick={() => setActiveView('chart')}
            >
              <BarChart className="h-4 w-4 mr-1" />
              Chart
            </Button>
          </div>
        </div>
      </CardHeader>
      
      <CardContent className="p-0">
        <div className="border-b">
          <button
            onClick={() => setShowQuery(!showQuery)}
            className="w-full flex items-center justify-between p-4 text-left hover:bg-gray-50"
          >
            <div className="flex items-center">
              <Code className="h-5 w-5 text-primary mr-2" />
              <span className="font-medium">SQL Query</span>
            </div>
            {showQuery ? (
              <ChevronUp className="h-5 w-5 text-gray-500" />
            ) : (
              <ChevronDown className="h-5 w-5 text-gray-500" />
            )}
          </button>
          
          {showQuery && (
            <div className="p-4 bg-gray-50 border-t">
              <pre className="bg-gray-900 text-gray-100 p-4 rounded-md overflow-x-auto">
                <code>{query}</code>
              </pre>
            </div>
          )}
        </div>
        
        <div className="p-4">
          <div className="flex items-center mb-4">
            <Database className="h-5 w-5 text-primary mr-2" />
            <h3 className="font-medium">Query Results</h3>
          </div>
          
          {activeView === 'table' ? (
            <div className="overflow-x-auto rounded-md border">
              <table className="w-full border-collapse">
                <thead>
                  <tr className="bg-gray-50">
                    {columns.map((column) => (
                      <th 
                        key={column} 
                        className="px-4 py-3 text-left text-sm font-medium text-gray-700 border-b"
                      >
                        {column}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {dataArray.length > 0 ? (
                    dataArray.map((row, rowIndex) => (
                      <tr 
                        key={rowIndex} 
                        className={rowIndex % 2 === 0 ? 'bg-white' : 'bg-gray-50'}
                      >
                        {columns.map((column) => (
                          <td 
                            key={`${rowIndex}-${column}`} 
                            className="px-4 py-3 text-sm text-gray-700 border-b"
                          >
                            {row[column]}
                          </td>
                        ))}
                      </tr>
                    ))
                  ) : (
                    <tr>
                      <td colSpan={columns.length || 1} className="px-4 py-3 text-sm text-gray-500 text-center">
                        No data available
                      </td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>
          ) : (
            <div className="h-64 flex items-center justify-center bg-gray-50 rounded-md border">
              <p className="text-gray-500">Chart visualization would be displayed here</p>
            </div>
          )}
          
          {data.length === 0 && (
            <div className="text-center py-8 text-gray-500">
              No results found
            </div>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
