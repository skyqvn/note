import java.sql.*;

//Connection： 数据库链接对象
//DriverManager： 驱动管理对
//Statement： sql语句的执行对
//ResultSet： DQL返回的查询结果集对
//SQLException： 数据库异常处理对象

public class mysqljc {
	public static String url = "jdbc:mysql://127.0.0.1/test";
	public static final String name = "com.mysql.cj.jdbc.Driver";
	public static final String user = "root";
	public static final String password = "123456";
	
	public static void main(String[] args) {
		Connection conn = null;
		Statement stmt = null;
		ResultSet rs;
		try {
			Class.forName(name);
			conn = DriverManager.getConnection(url, user, password);
			stmt = conn.createStatement();
			rs = stmt.executeQuery("select * from mytable2");
			while (rs.next()) {
				System.out.println(rs.getInt("id"));
				System.out.println(rs.getString("name"));
			}
		} catch (ClassNotFoundException | SQLException e) {
			e.printStackTrace();
		} finally {
			try {
				if (stmt != null) stmt.close();
			} catch (SQLException e) {
				e.printStackTrace();
			}
			
			try {
				if (conn != null) conn.close();
			} catch (SQLException e) {
				e.printStackTrace();
			}
		}
	}
}
