async function puanTablosunuGoster() {
    const response = await fetch("/puanlar");
    const puanlar = await response.json();

    const puanTablosuDiv = document.getElementById("puanTablosu");
    puanTablosuDiv.innerHTML = "<h3>Mevcut Puanlar:</h3>";

    for (const [oyuncu, puan] of Object.entries(puanlar)) {
        const p = document.createElement("p");
        p.textContent = `Oyuncu ${oyuncu}: ${puan} puan`;
        puanTablosuDiv.appendChild(p);
    }
}

async function puanGuncelle() {
    const oyuncuID = document.getElementById("oyuncuID").value;
    const puanDegisimi = document.getElementById("puanDegisimi").value;

    if (oyuncuID === "" || puanDegisimi === "") {
        alert("Lütfen tüm alanları doldurun.");
        return;
    }

    const response = await fetch(`/puanGuncelle?oyuncu=${oyuncuID}&puan=${puanDegisimi}`);
    if (response.ok) {
        alert("Puan başarıyla güncellendi.");
        puanTablosunuGoster();
    } else {
        alert("Hata: Oyuncu bulunamadı veya geçersiz puan değişimi.");
    }
}

document.addEventListener("DOMContentLoaded", puanTablosunuGoster);
