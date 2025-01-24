function App() {
	return (
		<div className="min-h-screen bg-gray-100">
			<div className="container mx-auto px-4 py-8">
				<div className="max-w-2xl mx-auto">
					<h1 className="text-3xl font-bold text-center mb-8">
						Secure File Sharing
					</h1>

					{/* File Upload Section */}
					<div className="bg-white rounded-lg shadow-md p-6">
						<div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
							<input type="file" className="hidden" id="fileInput" />
							<label
								htmlFor="fileInput"
								className="cursor-pointer inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
							>
								Choose File to Upload
							</label>
							<p className="mt-2 text-sm text-gray-600">
								or drag and drop your file here
							</p>
						</div>

						<div className="mt-4 text-sm text-gray-500">
							<p>Maximum file size: 10MB</p>
							<p>Supported formats: PDF, DOC, DOCX, TXT, PNG, JPG, JPEG, GIF</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}

export default App;
