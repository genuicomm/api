INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Contact',
    7,
    'contact',
    '[
      {
        "key": "title",
        "label": "Title",
        "type": "text",
        "value": "Hubungi Kami",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "label": "Sub Title",
        "type": "text",
        "value": "contact us",
        "is_visible": true
      },
      {
        "key": "description",
        "label": "Description",
        "type": "textarea",
        "value": "Kami siap membantu Anda mencapai potensi komunikasi terbaik",
        "is_visible": true
      },
      {
        "key": "address_office",
        "label": "Office Address",
        "type": "address",
        "value": "JI. Saturnus Timur XI No. 16 Kel. Manjahlega, Kec. Rancasari 40295 Kota Bandung",
        "link": "https://maps.google.com/?q=Jl. Saturnus Timur XI No. 16 Kel. Manjahlega, Kec. Rancasari Kota Bandung",
        "is_visible": true
      },
      {
        "key": "address_studio",
        "label": "Studio Address",
        "type": "address",
        "value": "JI. Pohon Cemara I No. 12 Komplek Garden City 2, Ciganitri-Buah Batu, Bandung 40287",
        "link": "https://maps.google.com/?q=Jl. Pohon Cemara I No. 12 Komplek Garden City 2, Ciganitri-Buah Batu, Bandung",
        "is_visible": true
      },
      {
        "key": "contact_info",
        "label": "Contact Info",
        "type": "object",
        "open_hours": "Senin - Jumat, 09.00 - 17.00",
        "email": "genuicomm@gmail.com",
        "whatsapp": "+62 877 6088 7862",
        "instagram": "@genuicomm.id",
        "facebook": "Genuicomm PS",
        "note": "",
        "is_visible": true
      }
    ]'::jsonb
);