# Deploy Your Code On a Docker Container Using Jenkins on AWS

![diagram-export-4-22-2025-7_09_24-PM](https://github.com/user-attachments/assets/44669afe-7ed3-46cd-b890-be4ec4b88dde)

Description 

## Agenda
1. Setup jenkins
2. Setup and configure git
3. Integration github with jenkins
4. Setup docker host
5. Integrate docker with jenkins
6. Automate the build and deploy process using jenkins
7. test the deployment

## Prequisites
1. Amazon Web Services (AWS)
2. SonarQube
3. Github account with the source code

## STEP 1 : Setup Jenkins Server on AWS EC2
1. Setup the linux instance
2. Install Java
3. Install Jenkins
4. Configure and Start Jenkins
5. Access web UI on PORT 8080

Login ke halaman Amazon Management Console, buka EC2 Dashboard, dan klik Launch Instance:
![Screenshot from 2025-04-22 19-20-46](https://github.com/user-attachments/assets/8702d2a9-4229-422e-accf-04ddc4cb8414)

Setelah meng-klik launch instance, berikan nama untuk server jenkins:
![Screenshot from 2025-04-22 19-31-18](https://github.com/user-attachments/assets/ad91c197-bcc9-4c93-a925-c1d1753be85a)

Selanjutnya memilih sistem operasi yang akan digunakan untuk server jenkins, disini saya memakai OS Ubuntu 22.04:
![Screenshot from 2025-04-22 19-39-41](https://github.com/user-attachments/assets/128d55a2-7fb7-41be-9441-144bade883d9)

Pilih instance type yang akan digunakan pada server jenkins:
![image](https://github.com/user-attachments/assets/07b7ff07-83b2-4e44-b21e-b77134ef4e09)

Selanjutnya membuat key pairs untuk mengakses server jenkins di terminal, buat key-pair baru:
![Screenshot from 2025-04-22 19-43-46](https://github.com/user-attachments/assets/cb999d63-c730-453f-bf8e-17c5a06e30ad)

Berikan nama untuk key-pair lalu pilih key pair 'RSA' dan file format nya '.pem'
![Screenshot from 2025-04-22 19-44-51](https://github.com/user-attachments/assets/c4418687-1df5-4c41-b23b-c5aeb78358b2)

Selanjutnya membuat storage, setelah itu klik launch instance:
![image](https://github.com/user-attachments/assets/dd64f31c-a3f0-43e7-97b7-1e8821fe7da3)

Buka terminal pada linux lalu ketikan perintah berikut ini untuk remote server aws ec2
```
ssh -i [file-ssh] [nama-user-server@ipaddress-server]
```
![Screenshot from 2025-05-02 16-16-44](https://github.com/user-attachments/assets/e1209936-06a1-4f20-ae08-a78f87d2eaf8)

Ubah nama server menjadi jenkins dengan cara mengetikkan perintah
```
sudo hostnamectl set-hostname jenkins
/bin/bash
```
![Screenshot from 2025-05-02 16-21-47](https://github.com/user-attachments/assets/b8327b33-dcdd-45f4-9166-068c566d0330)

Setelah merubaha nama server selanjutnya adalah menginstall java pada server jenkins, ketikkan perintah berikut ini untuk menginstall java di terminal
```
sudo apt install openjdk-17-jdk openjdk-17-jre

```
![Screenshot from 2025-04-30 14-40-26](https://github.com/user-attachments/assets/bc90eb53-4fd7-4248-9251-67f8798ad094)

Setelah berhasil install, ketikkan perintah berikut ini untuk memeriksa apakah penginstalan java sudah berhasil
![Screenshot from 2025-05-02 16-30-36](https://github.com/user-attachments/assets/362f2534-dbbc-4df4-8ced-106835d76209)

Setelah JAVA berhasil di install di server jenkins, selanjutnya adalah install jenkins di server. Kunjungsi website resmi jenkins lalu cari dokumentasi penginstalan jenkins pada OS Linux Ubuntu 
![Screenshot from 2025-05-02 16-35-30](https://github.com/user-attachments/assets/d368b169-3514-4f52-af8a-495c494d77e4)

Pilih perintah Long Term Support release, copy dan paste ke terminal server 
![Screenshot from 2025-04-30 14-48-54](https://github.com/user-attachments/assets/eee9930b-ed3c-4e2b-b719-8f1d270fd562)

Setelah proses instalasi jenkins selesai, selanjut nya aktifkan jenkins dan lihat status jenkins dengan perintah berikut ini 
```
sudo systemctl start jenkins
sudo systemctl status jenkins
```
![Screenshot from 2025-05-02 16-41-09](https://github.com/user-attachments/assets/8c360c42-93fe-4ff9-971c-3fb40e240143)

Sekarang akses jenkins dengan memasukkan ip server jengkins ditambah dengan port default nya yaitu 8080
```
52.77.231.48:8080
```
![Screenshot from 2025-04-30 14-51-23](https://github.com/user-attachments/assets/39b490a8-0359-43c7-8d29-f4fc59b67492)

Customize halaman jenkins, pilih install yang di rekomendasikan

![customize jenkins](https://github.com/user-attachments/assets/8bbf21d9-eb0c-4b95-b13b-d6b917ac1715)

Buat buat akun admin user dan isi kolom yang ada pada gambar berikut

![WhatsApp Image 2025-05-02 at 17 36 57](https://github.com/user-attachments/assets/32723058-f463-4deb-86e7-cd9de7412013)





