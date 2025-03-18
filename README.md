# gPRC
## Định nghĩa
gRPC (gRPC Remote Procedure Call) là một framework mã nguồn mở được phát triển bởi Google, cho phép các dịch vụ giao tiếp với nhau một cách hiệu quả, nhanh chóng và đa ngôn ngữ.

## Đặc điểm chính của gRPC
* Sử dụng giao thức HTTP/2, cho phép truyền dữ liệu song song và giảm độ trễ.
* Sử dụng Protocol Buffers (protobuf) làm định dạng dữ liệu mặc định, giúp giảm kích thước dữ liệu và tăng tốc độ truyền.
* Hỗ trợ nhiều ngôn ngữ lập trình như Go, Java, Python, C++, Node.js, ...
* Khả năng tự động sinh code dựa trên proto do người dùng định nghĩa.
* Hỗ trợ các mô hình giao tiếp:
	+ Unary: 1 request, 1 response
	+ Server Streaming: 1 request, nhiều response
	+ Client Streaming: nhiều request, 1 response
	+ Bidirectional Streaming: nhiều request, nhiều response

## Ưu điểm của gRPC
* Hiệu suất cao, tối ưu cho hệ thống phân tán.
* Hỗ trợ load balancing và health checking.
* Giảm độ trễ nhờ HTTP/2 và protobuf.
* Dễ dàng mở rộng trong hệ thống microservices.

## Nhược điểm của gRPC
* Khó khăn trong việc debug lỗi.
* Số lượng browser hỗ trợ hạn chế do sử dụng HTTP/2.
* Khó khăn trong việc tiếp cận hơn so với REST.
* Ít công cụ hỗ trợ hơn so với REST.

## So sánh gRPC và REST
| | gRPC | REST |
| -------- | ------- | ------ |
| Phương pháp thiết kế | Thiết kế hướng dịch vụ | Thiết kế hướng thực thể |
| Mô hình truyền thông | Nhiều lựa chọn (1 hoặc nhiều clients giao tiếp với 1 hoặc nhiều server) | Chỉ 1 client giao tiếp với 1 server |
| Dữ liệu trả về | Cố định kiểu dữ liệu trả về như định nghĩa trên proto | Cố định theo structure (JSON) định nghĩa trên server |
| Liên kết giữa client và server | Client và server cần proto giống nhau để định nghĩa format của dữ liệu | Client và server không cần biết rõ về nhau |
| Use case | Kiến trúc microservices cần hiệu suất cao hoặc trả về nhiều dữ liệu | Nguồn dữ liệu đơn giản, trong đó tài nguyên được xác định rõ ràng |