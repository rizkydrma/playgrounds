const fs = require('fs');

// https://api.blabla.com/
const baseUrl = '';
const perPage = 1000;
const outputFilename = 'indonesia_schools.csv';

function saveToCSV(data, filename) {
	// Jika file belum ada, tulis header
	if (!fs.existsSync(filename)) {
		const headers = Object.keys(data[0]).join(',');
		fs.writeFileSync(filename, headers + '\n');
	}

	// Fungsi untuk mengatasi koma dalam data
	const escapeCommas = (str) => {
		// Jika string mengandung koma, bungkus dalam tanda kutip ganda
		if (str.includes(',')) {
			return `"${str}"`;
		}
		return str;
	};

	// Convert each data item to CSV format
	const csvData = data
		.map((item) => {
			// Map each value in the item to a comma-separated string,
			// escaping commas if necessary
			return Object.values(item).map(escapeCommas).join(',');
		})
		.join('\n');

	// Append the CSV data to the file
	fs.appendFileSync(filename, csvData + '\n');
}

async function fetchData() {
	let currentPage = 1;
	let totalData = 0;

	// First fetch to know total data
	const initialResponse = await fetch(`${baseUrl}?page=${currentPage}&perPage=1`);
	const initialData = await initialResponse.json();
	totalData = initialData.total_data;

	// Calculate total page
	const totalPages = Math.ceil(totalData / perPage);

	// Loop melalui setiap halaman dan ambil data
	for (currentPage = 1; currentPage <= totalPages; currentPage++) {
		const response = await fetch(`${baseUrl}?page=${currentPage}&perPage=${perPage}`);
		const data = await response.json();
		const dataSekolah = data.dataSekolah;

		// Simpan data ke CSV
		saveToCSV(dataSekolah, outputFilename);

		console.log(`Data dari halaman ${currentPage} disimpan.`);
	}
}

fetchData()
	.then(() => {
		console.log('Semua data berhasil diambil dan disimpan ke file CSV.');
	})
	.catch((error) => {
		console.error('Terjadi kesalahan saat mengambil data:', error);
	});
